package service

import (
	"os"

	signal "github.com/multiverse-os/service/signal"
)

// TODO: This is a clever way to expose our object type to add methods
//type SignalHandler signal.Handler

var ShutdownSignals = signal.ShutdownSignals

func NewHandler() signal.Handler { return signal.NewHandler() }

func OnShutdownSignals(function func(os.Signal)) signal.Handler {
	return signal.ShutdownHandler(function)
}

func ShutdownHandler(function func(os.Signal)) signal.Handler {
	return signal.ShutdownHandler(function)
}
