package daemon

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"
)

// [ Test Function ]////////////////////////////////////////////////////////////
func Daemonize(function func()) error {
	fmt.Println("Daemonizing")

	daemon := Process{LogFile: "/dev/stdout"}
	if child, err := daemon.Run(); err != nil {
		return err
	} else {
		if child != nil {
			return fmt.Errorf("[error] failed to create daemon process")
		}
	}

	// TODO: This would setup the pid file previously, should require user to
	// supply the PID file location or default to system default
	//defer daemon.Release()
	function()
	return nil
}

///////////////////////////////////////////////////////////////////////////////

type EnvVar struct {
	Key   string
	Value string
}

// TODO: Not a huge fan of having 5 versions of Process
type Process struct {
	EnvVar
	EnvVars          map[string]string
	Initialized      bool
	Pid              int
	PidFile          string
	LogDirectory     string
	LogFile          string
	WorkingDirectory string
	Chroot           string
	Env              []string
	Args             []string
	Credential       *syscall.Credential
	Umask            int
	abspath          string
	logFile          *os.File
	nullFile         *os.File
	rpipe, wpipe     *os.File
}

func (self *Process) Environment() (env []string) {
	for envKey, envValue := range self.EnvVars {
		env = append(env, fmt.Sprintf("%s=%s", envKey, envValue))
	}
	return env
}

func (self *Process) AppendEnvVar(envKey, envValue string) *Process {
	self.EnvVars[envKey] = envValue
	return self

}

func (self *Process) ParseEnvironment() (env map[string]string) {
	rawEnv := os.Environ()
	for _, rawVar := range rawEnv {
		envVar := strings.Split(rawVar, "=")
		env[envVar[0]] = envVar[1]
	}
	self.EnvVars = env
	return env
}

func (self *Process) IsRunning() bool {
	// TODO: Is there really no better way to handle this?
	return os.Getenv("_DAEMON") == "1"
}

func (self *Process) Run() (childProcess *os.Process, err error) {
	fmt.Println("Run() command")
	if self.IsRunning() {
		if err = self.child(); err != nil {
			return nil, err
		}
	} else {
		if childProcess, err = self.parent(); err != nil {
			return nil, err
		}
	}
	return childProcess, nil
}

func (self *Process) parent() (parentProcess *os.Process, err error) {
	if err = self.prepareEnv(); err != nil {
		return
	}
	attr := &os.ProcAttr{
		Dir:   self.WorkingDirectory,
		Env:   self.Env,
		Files: self.files(),
		Sys: &syscall.SysProcAttr{
			//Chroot:     self.Chroot, // To bad there is no comments on this
			Credential: self.Credential,
			Setsid:     true,
		},
	}
	if parentProcess, err = os.StartProcess(self.abspath, self.Args, attr); err != nil {
		// TODO: Replace this with our pid code
		//if self.pidFile != nil {
		//	self.pidFile.Remove()
		//}
		return
	}
	self.rpipe.Close()
	encoder := json.NewEncoder(self.wpipe)
	if err = encoder.Encode(self); err != nil {
		return
	}
	_, err = fmt.Fprint(self.wpipe, "\n\n")
	return
}

func (self *Process) prepareEnv() (err error) {
	if self.abspath, err = os.Executable(); err != nil {
		return err
	}
	if len(self.Args) == 0 {
		self.Args = os.Args
	}
	mark := fmt.Sprintf("%s=%s", self.EnvVar.Key, self.EnvVar.Value)
	if len(self.Env) == 0 {
		self.Env = os.Environ()
	}
	self.Env = append(self.Env, mark)
	return nil
}

func (self *Process) files() (processFiles []*os.File) {
	log := self.nullFile
	if self.logFile != nil {
		log = self.logFile
	}
	// TODO: Default should be nil
	processFiles = []*os.File{
		self.rpipe,    // (0) stdin
		log,           // (1) stdout
		log,           // (2) stderr
		self.nullFile, // (3) dup on fd 0 after initialization
	}
	// TODO: PID file defintion was here.
	//if self.pidFile != nil {
	//	f = append(f, self.pidFile.File) // (4) pid file
	//}

	// TODO: This is the method using our pid system
	//pid.Write(self.PIDFile)

	return processFiles
}

func (self *Process) child() error {
	if self.Initialized {
		return os.ErrInvalid
	}
	self.Initialized = true
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(self); err != nil {
		//pid.Clean(self.PIDFile)
		return err
	}
	// create PID file after context decoding to know PID file full path.
	// TODO: Previously file locking was used, but this may not actually be
	// necessary unless we scale this up to handle several processes
	// TODO: we should just have a cleanup function and a encapsulating function
	// that will run it if error, instead of calling pid.Clean more than once.
	//pid.Write(self.PIDFile)
	if err := syscall.Close(0); err != nil {
		//pid.Clean(self.PIDFile)
		return err
	}
	if err := syscallDup(3, 0); err != nil {
		//pid.Clean(self.PIDFile)
		return err
	}
	if self.Umask != 0 {
		syscall.Umask(int(self.Umask))
	}
	if len(self.Chroot) > 0 {
		if err := syscall.Chroot(self.Chroot); err != nil {
			//pid.Clean(self.PIDFile)
			return err
		}
	}
	return nil
}
