package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Format uint8

const (
	YAML Format = iota
	JSON
	TOML
	INI
)

type File struct {
	Path   string
	Format string
}

type Maintainance struct {
	Enabled      bool   `yaml:"enabled"`
	Sessions     bool   `yaml:"sessions"`
	Announcement string `yaml:"announcement"`
}

type Directories struct {
	Data  string `yaml:"data"`
	Tmp   string `yaml:"tmp"`
	Cache string `yaml:"cache"`
}

// TODO: Would prefer it to all be under app: like seen in rails,  and this can
// be done using our own marshal and unmarshal funcitons would should be done
// regardless for maximum control
// TODO: Address/Port should be handled in an nginx like configuration since
// this application framework is meant to be able to handle reverse proxy,
// multiple hosts/domains
// TODO: maybe directories should be in their owns ection so its like
// directories:
//
//	data: ".."
//	cache:
//	tmp:
//
// etc..
type Settings struct {
	Name         string       `yaml:"name"`
	Environment  string       `yaml:"environment"`
	Address      string       `yaml:"address"`
	Port         int          `yaml:"port"`
	PidFilename  string       `yaml:"pid"`
	Directories  Directories  `yaml:"directories"`
	Maintainance Maintainance `yaml:"maintainance"`
}

// TODO: Add ability to update a setting. Write a default settings file.
// TODO: Separate out application specific settings from config library logic so
// we can re-use this code.

func Load(path string) (config *Settings, err error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func (ss Settings) InitializeDefaultConfig(path string) error {
	return DefaultConfig("app").Save(path)
}

func (ss Settings) Save(path string) error {
	configPath, _ := filepath.Split(path)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.MkdirAll(configPath, 0700)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := ss.Save(path)
		if err != nil {
			return nil
		}
	} else {
		return err
	}
	return nil
}

// TODO: We need to also parse out environmental variables
// TODO: Using localhost opens up to DNS rebind attacks
func DefaultConfig(app string) *Settings {
	return &Settings{
		Name:        "app",
		Environment: "development",
		Address:     "localhost",
		Port:        3000,
		PidFilename: app + ".pid",
		Directories: Directories{
			Data:  "~/.local/share/maglev/" + app,
			Cache: "~/.cache/maglev/" + app,
		},
		Maintainance: Maintainance{
			Enabled:      false,
			Sessions:     false,
			Announcement: "Updating Server!",
		},
	}
}

func Validate(config *Settings) *Settings {
	if len(config.Name) == 0 {
		config.Name = DefaultConfig("app").Name
	}
	// TODO: Need more validations for all the individual fields
	// TODO: Port needs to only support actual ports 1 - ~65000
	if len(config.PidFilename) == 0 {
		config.PidFilename = "app/tmp/pids/" + config.Name + ".pid"
	}
	if len(config.Directories.Data) == 0 {
		config.Directories.Data = "~/.local/share/maglev/" + config.Name
	}
	if len(config.Directories.Cache) == 0 {
		config.Directories.Cache = "~/.cache/maglev/" + config.Name
	}
	return config
}
