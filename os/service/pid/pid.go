package pid

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
)

type File struct {
	*os.File
	Path string
	Pid  int
}

func New(pidPath string) *File {
	pidPath = ValidatePath(pidPath)
	return &File{
		Path: pidPath,
		Pid:  os.Getpid(),
	}
}

func (f *File) Clean() error {
	unlock(f.Fd())
	f.Close()
	return removeFile(f.Path)
}

func OSDefault() string {
	app := serviceName()
	return ("/var/run/" + app + "/" + app + ".pid")
}
func TempDefault() string {
	return ("/var/tmp/" + serviceName() + ".pid")
}
func UserDefault() string {
	user, err := user.Current()
	if err != nil {
		return TempDefault()
	}
	return ("/run/" + user.Name + "/" + serviceName() + ".pid")
}

func WriteToTempDirectory() (*File, error) { return Write(TempDefault()) }
func WriteToOSDefault() (*File, error)     { return Write(OSDefault()) }
func WriteToUserDefault() (*File, error)   { return Write(UserDefault()) }

// TODO: Consider using os.Executable() as the default app name
func ValidatePath(pidPath string) string {
	if len(pidPath) < 0 || len(pidPath) > 256 {
		return OSDefault()
	}
	basename := path.Base(pidPath)
	if basename[len(basename)-1:] == "/" {
		fmt.Printf("writing pid to: %v.pid", (pidPath + serviceName()))
		return pidPath + serviceName() + ".pid"
	} else if filepath.Ext(basename) != ".pid" {
		fmt.Printf("writing pid file: %v.pid", pidPath)
		return pidPath + ".pid"
	} else {
		return pidPath
	}
}

// TODO
// Test this; it may not work if locked, so below fileless clean and I'm
// pretty sure there is a better way to avoid failure
func removeFile(path string) error {
	if err := os.Remove(path); err != nil {
		return errCleanFailed
	}
	return nil
}

func serviceName() string {
	executable, _ := os.Executable()
	return path.Base(executable)
}

func Write(pidPath string) (*File, error) {
	pid := New(pidPath)
	// NOTE: Confirm path exists, if does not exist write it
	directory := filepath.Dir(pid.Path)
	if _, err := os.Stat(pid.Path); os.IsNotExist(err) {
		if err := os.MkdirAll(directory, 0700); err != nil {
			Write(TempDefault())
		}
	}
	if _, err := os.Stat(pid.Path); !os.IsNotExist(err) {
		// NOTE: Exists, checking if pid is stale
		if pid.File, err = os.OpenFile(pid.Path, os.O_RDWR|os.O_CREATE, 0600); err != nil {
			return nil, errStalePid
		} else {
			if pidData, err := ioutil.ReadFile(pid.Path); err != nil {
				pidInt, _ := strconv.Atoi(string(pidData))
				if isProcessRunning(pidInt) {
					return nil, errFileLocked
				} else {
					Clean(pid.Path)
				}
			}
		}
	}
	// NOTE: Standard creation, file locking, and return Pid file object
	var err error
	if pid.File, err = os.OpenFile(pid.Path, os.O_RDWR|os.O_CREATE, 0600); err != nil {
		return nil, errOpenFailed
	} else {
		if _, err := pid.File.WriteString(strconv.Itoa(pid.Pid) + "\n"); err != nil {
			return nil, errWriteFailed
		}
	}
	// NOTE: Locking via Fd() and returning the File object
	lock(pid.File.Fd())
	return pid, nil
}

func Clean(pidPath string) error {
	if _, err := os.Stat(pidPath); os.IsNotExist(err) {
		return nil
	} else {
		// TODO: Experimental way to try to close it out with just path
		// this was written without actually testing since we moved to holding the
		// file in memory
		if err := removeFile(pidPath); err != nil {
			file, err := os.OpenFile(pidPath, os.O_RDWR, 0600)
			if err != nil {
				return nil
			}
			unlock(file.Fd())
			return removeFile(pidPath)
		}
	}
	return nil
}
