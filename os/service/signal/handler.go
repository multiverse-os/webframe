package signal

import (
	"os"
	"os/signal"
	"sync"
)

type Handler struct {
	hooks     map[os.Signal][]func(os.Signal)
	ignored   []os.Signal
	mutex     *sync.Mutex
	channel   chan os.Signal
	listening bool
}

// TODO: Should it be optional to declare the channel handling here so the
// design where one uses signal handling to hold open the application?
func NewHandler() Handler {
	handler := Handler{
		hooks:   map[os.Signal][]func(os.Signal){},
		ignored: []os.Signal{},
		mutex:   &sync.Mutex{},
		channel: make(chan os.Signal, 1),
	}
	go func() {
		for {
			incomingSignal := <-handler.channel
			handler.handle(incomingSignal)
		}
	}()
	return handler
}

func ShutdownHandler(function func(os.Signal)) Handler {
	handler := NewHandler()
	return handler.OnShutdown(function)
}

func (self Handler) handle(incomingSignal os.Signal) {
	functions := self.hooks[incomingSignal]
	for _, function := range functions {
		function(incomingSignal)
	}
}

// /////////////////////////////////////////////////////////////////////////////
func (h Handler) Add(function func(os.Signal), signals ...os.Signal) Handler {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	for _, s := range signals {
		h.hooks[s] = append(h.hooks[s], function)
		signal.Notify(h.channel, s)
	}
	return h
}

// NOTE: Shutdown adds the defined hook to all shutdown/exit signals
func (h Handler) OnShutdown(function func(os.Signal)) Handler {
	return h.Add(function, ShutdownSignals...)
}

func (h Handler) OnInterrupt(function func(os.Signal)) Handler {
	return h.Add(function, Interrupt)
}
func (h Handler) OnTerminate(function func(os.Signal)) Handler {
	return h.Add(function, Terminate)
}
func (h Handler) OnQuit(function func(os.Signal)) Handler   { return h.Add(function, Quit) }
func (h Handler) OnHangup(function func(os.Signal)) Handler { return h.Add(function, Hangup) }
func (h Handler) OnKill(function func(os.Signal)) Handler   { return h.Add(function, Kill) }

// /////////////////////////////////////////////////////////////////////////////
// TODO: Would be nice to eventually build in functionality to turly ignore
// signals even force kill signals; possibly via uninterruptable sleep or some
// similar technique to make this a truly general use signal handling library
func (h Handler) Ignore(signals ...os.Signal) Handler {
	h.ignored = append(h.ignored, signals...)
	for _, ignoredSignal := range h.ignored {
		h.Remove(ignoredSignal)
		signal.Ignore(h.Signals()...)
	}
	return h
}

// NOTE: Stop ignoring will stop listening too since it requires use of reset
// which resets both Notify and Ignore calls
func (h Handler) StopIgnoring() Handler {
	signal.Reset()
	return h
}

func (h Handler) StopListening() Handler {
	signal.Stop(h.channel)
	return h
}

// /////////////////////////////////////////////////////////////////////////////
func (h Handler) Remove(s os.Signal) Handler {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	delete(h.hooks, s)
	return h
}

func (h Handler) Clear() Handler {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	for s := range h.hooks {
		delete(h.hooks, s)
	}
	return h
}

func (h Handler) Reset() Handler {
	h.Clear()
	h.StopListening()
	return h
}

// /////////////////////////////////////////////////////////////////////////////
func (h Handler) Signals() []os.Signal {
	signals := make([]os.Signal, len(h.hooks))
	for s, _ := range h.hooks {
		signals = append(signals, s)
	}
	return signals
}
