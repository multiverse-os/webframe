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
type AppDirectories struct {
	Working    string
	Data       string
	Temporary  string
	UserHome   string
	UserCache  string
	UserConfig string
	UserData   string
}

func (self AppDirectories) ParseApplicationDirectories() {
	var err error
	if self.Working, err = os.Getwd(); err != nil {
		panic(fmt.Sprintf("failed to determine working directory:", err))
	}
	self.Temporary = os.TempDir()
}

func (self *AppDirectories) ParseUserDirectories() {
	var err error
	self.UserHome = os.Getenv("HOME")
	// TODO: Why is this undefined?
	// REF: https://golang.org/src/os/file.go
	//self.UserHomeDirectory, err = os.UserHomeDir()
	//if err != nil {
	//	panic(fmt.Sprintf("[fatal error] failed to determine user home:", err))
	//}
	if self.UserCache, err = os.UserCacheDir(); err != nil {
		panic(fmt.Sprintf("failed to determine user cache:", err))
	}
	if self.UserConfig = self.UserHome + "/.config/starship"; err != nil {
		panic(fmt.Sprintf("failed to determine user config path:", err))
	}
	if _, err := os.Stat(self.UserConfig); os.IsNotExist(err) {
		os.Mkdir(self.UserConfig, 0770)
	}
	if self.UserData = self.UserHome + "/.local/share/starship/"; err != nil {
		panic(fmt.Sprintf("failed to determine user data path:", err))
	}
	if _, err := os.Stat(self.UserData); os.IsNotExist(err) {
		os.Mkdir(self.UserData, 0770)
	}
}
