package signal

import (
	"os"
	"syscall"
)

// User Defined Signal Handlers
// A process can replace the default signal handler for almost all signals (but not SIGKILL) by its user’s own handler function.
// A signal handler function can have any name, but must have return type void and have one int parameter.
// Example: you might choose the name sigchld_handler for a signal handler for the SIGCHLD signal (termination of a child process). Then the declaration would be:
//
//	void sigchld_handler(int sig);
//type SignalType int
//
//const (
//	ShutdownType SignalType = iota
//)

var (
	// NOTE: This allows packages using this signal system to have a signal import
	// related to signals and avoid an extra unnecessary syscall import.
	Interrupt os.Signal = syscall.SIGINT
	Terminate os.Signal = syscall.SIGTERM
	Quit      os.Signal = syscall.SIGQUIT
	Hangup    os.Signal = syscall.SIGHUP
	Kill      os.Signal = syscall.SIGKILL
	// Aliasing for a more intuitive API
	SIGINT  = Interrupt
	SIGTERM = Terminate
	SIGQUIT = Quit
	SIGHUP  = Hangup
	SIGKILL = Kill
	//Ign: Ignore the signal; i.e., do nothing, just return
	//Stop: block the process
)

// Signals
// /////////////////////////////////////////////////////////////////////////////
// REF: runtime/sigtab_linux_generic.go (Go source code)
// These are what is defined below as the shutdown or exit signal group
// SIGHUP: terminal line hangup
// SIGINT: interrupt
// SIGQUIT: quit (core dump expected)
// SIGABRT: abort
// SIGTERM: termination
var ShutdownSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGHUP,
	syscall.SIGKILL,
}
