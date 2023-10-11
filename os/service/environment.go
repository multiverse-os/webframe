package service

import (
	"os"
	"strings"
)

func ParseEnvironment() (env map[string]string) {
	rawEnv := os.Environ()
	for _, rawVar := range rawEnv {
		envVar := strings.Split(rawVar, "=")
		env[envVar[0]] = envVar[1]
	}
	return env
}
