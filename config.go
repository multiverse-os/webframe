package maglev

import (
	config "github.com/multiverse-os/maglev/config"
)

func DefaultConfig() config.Settings { return config.DefaultConfig() }

func LoadConfig(path string) (config.Settings, error) {
	return config.Load(path)
}
