package service

import (
	pid "github.com/multiverse-os/service/pid"
)

func WritePid(path string) { pid.Write(path) }
func CleanPid(path string) { pid.Clean(path) }
