package path

import (
	"os"
	"path/filepath"
	"strings"
)

// propagate app configuration.
type App struct {
	Name string
}

func (a App) path(envVar string, defaultFn func() string, file string) string {
	base := os.Getenv(envVar)
	if len(base) == 0 {
		base = defaultFn()
	}
	return filepath.Join(base, a.Name, file)
}

func (a App) DataPath(file string) string   { return a.path("XDG_DATA_HOME", DataHome, file) }
func (a App) ConfigPath(file string) string { return a.path("XDG_CONFIG_HOME", ConfigHome, file) }
func (a App) CachePath(file string) string  { return a.path("XDG_CACHE_HOME", CacheHome, file) }

func (a App) SystemDataPaths(file string) []string {
	return a.multiPaths("XDG_DATA_DIRS", DataDirs, file)
}

func (a App) SystemConfigPaths(file string) []string {
	return a.multiPaths("XDG_CONFIG_DIRS", ConfigDirs, file)
}

// TODO: This is clearly wrong, why append to dirs then return directories??
func (a App) multiPaths(envVar string, defaultFn func() []string, file string) (directories []string) {
	env := os.Getenv(envVar)
	if len(env) != 0 {
		directories = strings.Split(env, ":")
	} else {
		directories = defaultFn()
	}
	var dirs []string
	for _, it := range directories {
		dirs = append(dirs, filepath.Join(it, a.Name, file))
	}
	return dirs
	//return directories
}

func LoadApplicationData(name string) App {
	return App{
		Name: name,
	}
}

func DataHome() string {
	return os.Getenv("HOME") + "/.local/share"
}

func ConfigHome() string {
	return os.Getenv("HOME") + "/.config"
}

func CacheHome() string {
	return os.Getenv("HOME") + "/.cache"
}

func DataDirs() []string {
	// The specification gives a  value with trailing slashes but only
	// for this value. Seems odd enough to take the liberty of removing them.
	return []string{"/usr/local/share", "/usr/share"}
}

func ConfigDirs() []string {
	return []string{"/etc/xdg"}
}
