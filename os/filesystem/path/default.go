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

func (self App) path(envVar string, defaultFn func() string, file string) string {
	base := os.Getenv(envVar)
	if len(base) == 0 {
		base = defaultFn()
	}
	return filepath.Join(base, self.Name, file)
}

func (self App) DataPath(file string) string   { return self.path("XDG_DATA_HOME", DataHome, file) }
func (self App) ConfigPath(file string) string { return self.path("XDG_CONFIG_HOME", ConfigHome, file) }
func (self App) CachePath(file string) string  { return self.path("XDG_CACHE_HOME", CacheHome, file) }

func (self App) SystemDataPaths(file string) []string {
	return self.multiPaths("XDG_DATA_DIRS", DataDirs, file)
}

func (self App) SystemConfigPaths(file string) []string {
	return self.multiPaths("XDG_CONFIG_DIRS", ConfigDirs, file)
}

func (a App) multiPaths(envVar string, defaultFn func() []string, file string) (directories []string) {
	env := os.Getenv(envVar)
	if len(env) != 0 {
		directories = strings.Split(env, ":")
	} else {
		directories = defaultFn()
	}
	var dirs []string
	for _, it := range directores {
		dirs = append(dirs, filepath.Join(it, a.Name, file))
	}
	return directories
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
