package filesystem

import (
	"fmt"
	"os"
)

// TODO: Id prefer to call it like Global vs Local and we should have a more
// generic directories type that stores weather its global vs local to avoid
// repeating fields (diy); and possibly even more generic to define collections
// of folders for a purpose-- this is specifically for application directories
// for accessing the disk in the standardized and expected way
type Directories struct {
	App        string
	Working    string
	Data       string
	Temporary  string
	UserHome   string
	UserCache  string
	UserConfig string
	UserData   string
}

func (d Directories) ParseApplicationDirectories() {
	var err error
	if d.Working, err = os.Getwd(); err != nil {
		panic(fmt.Sprintf("failed to determine working directory:", err))
	}
	d.Temporary = os.TempDir()
}

func (d *Directories) ParseUserDirectories() {
	var err error
	d.UserHome = os.Getenv("HOME")
	// TODO: Why is this undefined?
	// REF: https://golang.org/src/os/file.go
	//self.UserHomeDirectory, err = os.UserHomeDir()
	//if err != nil {
	//	panic(fmt.Sprintf("[fatal error] failed to determine user home:", err))
	//}
	if d.UserCache, err = os.UserCacheDir(); err != nil {
		panic(fmt.Sprintf("failed to determine user cache:", err))
	}
	if d.UserConfig = d.UserHome + "/.config/" + d.App; err != nil {
		panic(fmt.Sprintf("failed to determine user config path:", err))
	}
	if _, err := os.Stat(d.UserConfig); os.IsNotExist(err) {
		os.Mkdir(d.UserConfig, 0770)
	}
	if d.UserData = d.UserHome + "/.local/share/" + d.App + "/"; err != nil {
		panic(fmt.Sprintf("failed to determine user data path:", err))
	}
	if _, err := os.Stat(d.UserData); os.IsNotExist(err) {
		os.Mkdir(d.UserData, 0770)
	}
}
