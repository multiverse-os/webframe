package webframe

import (
	"fmt"
	"os"

	service "github.com/multiverse-os/service"
)

// TODO: Could use a map[signal.Signal]func() error and then we can allow the
// app to

func (f *Framework) InitSignalHandler() {
	f.Process.Signals = service.OnShutdownSignals(func(shutdownSignal os.Signal) {
		switch shutdownSignal {
		case os.Interrupt:
			fmt.Printf("\n")
		}
		f.Log(fmt.Sprintf("received exit signal: %v", shutdownSignal))
		f.Stop()
	})
}

func (f Framework) WritePid() { f.Process.WritePid(f.Config.PidFilename) }
func (f Framework) CleanPid() { f.Process.CleanPid() }

func (f Framework) Start() {
	f.WritePid()
	// TODO: Create a servers.Start() to do them all together more easily?
	for _, server := range f.servers {
		server.Start()
	}

	// TODO: This should not be an empty open, we should be starting a mainLoop
	// style function that can handle our job firing to begin with.
	for {
	}
}

func (f Framework) ParseEnvConfig() {
	env := service.ParseEnvironment() // map[string]string of values
	// TODO: Build a config.Settings or just assign to current framework object (i
	// prefer the first option)
	f.Log("env=(%v)", env)
	// TODO: Then use this in the cmd tool to do proper config variable loading
	// cascade
}

func (f Framework) Stop() {
	f.Log("initiating cleanup sequence, and stopping the application")
	f.ShutdownFunctions()
	f.CleanPid()
	os.Exit(0)
}

func (f Framework) Restart() {
	f.Log("application is attempting to restart")
	for _, server := range f.servers {
		server.Restart()
	}
}

func (f *Framework) AppendToShutdown(exitFunction func()) {
	f.shutdown = append(f.shutdown, exitFunction)
}

func (f Framework) ShutdownFunctions() {
	// TODO: Using a Debug boolean, we should not encapsulate each of these which
	// would become incredibly unwieldy; instead we should have a check inside
	// f.Dubug(text) that does the bool Debug check in framework so that these can
	// be called like this and can remain if useful without change
	f.Debug(fmt.Sprintf("number of shutdown functions(%v)", len(f.shutdown)))
	for _, function := range f.shutdown {
		function()
	}
}
