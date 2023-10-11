package main

import (
	"fmt"
	"os"

	service "github.com/multiverse-os/service"
)

// TODO: Laying out the foundation for an example function for dealing with
// services this is not unique at all to the web framework so it makes far more
// sense

// TODO: Probably should throw more stuff like startup and shutdown and pid path
// into an object we can call into any project

type Service struct {
	Process *service.Process
}

func main() {
	fmt.Println("service-cli")
	fmt.Println("===========")
	s := Service{
		Process: service.ParseProcess(),
	}

	fmt.Printf("parsed process(%v)\n", s)

	s.InitSignalHandler()
	s.Start()
}

func (s *Service) InitSignalHandler() {
	s.Process.Signals = service.OnShutdownSignals(func(shutdownSignal os.Signal) {
		switch shutdownSignal {
		case os.Interrupt:
			fmt.Printf("\n")
		}
		fmt.Printf(fmt.Sprintf("received exit signal: %v\n", shutdownSignal))
		s.Stop()
	})
}

func (s *Service) Start() {
	service.WritePid(".service-test")

	// TODO: This should not be an empty open, we should be starting a mainLoop
	// style function that can handle our job firing to begin with.
	for {
	}
}

// TODO
// This function is replicated in fork, same with IsProcessRunning()
func (s *Service) ParseEnvConfig() {
	env := service.ParseEnvironment() // map[string]string of values
	// TODO: Build a config.Settings or just assign to current framework object (i
	// prefer the first option)
	fmt.Printf("env=(%v)\n", env)
	// TODO: Then use this in the cmd tool to do proper config variable loading
	// cascade
}

func (s *Service) Stop() {
	fmt.Printf("initiating cleanup sequence, and stopping the application\n")
	s.Process.Shutdown()
	service.CleanPid("~/.service-test")
	os.Exit(0)
}

func (s *Service) ShutdownFunctions() {
	// TODO: Using a Debug boolean, we should not encapsulate each of these which
	// would become incredibly unwieldy; instead we should have a check inside
	// f.Dubug(text) that does the bool Debug check in framework so that these can
	// be called like this and can remain if useful without change
	fmt.Printf("shutdown functions count(%v)\n", s.Process.ShutdownCount())
	s.Process.Shutdown()
}
