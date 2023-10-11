package service

import (
	daemon "github.com/multiverse-os/service/daemon"
)

// TODO: Abstract the logic from each of the dependency free subpackages so that
// a developer can access all the necessary functionality from each package to
// setup a quality Linux service in an intuitive way and doing this only
// importing the base service package. Even if functions are essentially
// duplicated, this is fine for a simple interface to 5 subpackages

func Daemonize(function func()) error {
	return daemon.Daemonize(function)
}
