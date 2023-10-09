package event

import (
	"fmt"
	"sync"
	"time"
)

// Listener is the function type to run on events.
type Listener func(interface{})

// TODO: Cant quit be an event?
type Watcher struct {
	//name           string
	quit           chan bool
	events         chan interface{}
	listeners      []Listener
	mutex          *sync.Mutex
	bufferEvents   []interface{}
	bufferDuration time.Duration
	Verbose        bool
}

func Watch() *Watcher {
	watcher := &Watcher{}
	watcher.Start()
	return watcher
}

// Open the observer channles and run the event loop,
// it will return an error if event loop already running.
func (w *Watcher) Start() error {
	// Check for mutex
	if w.mutex == nil {
		w.mutex = &sync.Mutex{}
	}

	if w.events != nil {
		return fmt.Errorf("Watcher already inititated.")
	}

	// Create the observer channels.
	w.quit = make(chan bool)
	w.events = make(chan interface{})

	// Run the observer.
	return w.eventLoop()
}

// TODO: Would it be anymore efficient to have a single channel, at the very
// least it may be simpler; and easier to understand and therefore customize
// Close the observer channles,
// it will return an error if close fails.
func (w *Watcher) Stop() error {
	// Close event loop
	if w.events != nil {
		// Send a quit signal.
		w.quit <- true

		// Close channels.
		close(w.quit)
		close(w.events)
	}

	return nil
}

// AddListener adds a listener function to run on event,
// the listener function will recive the event object as argument.
func (w *Watcher) AddListener(l Listener) {
	// Check for mutex
	if w.mutex == nil {
		w.mutex = &sync.Mutex{}
	}

	// Lock:
	// 1. operations on array listeners
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.listeners = append(w.listeners, l)
}

// Emit an event, and event can be of any type, when event is triggered all
// listeners will be called using the event object.
func (w *Watcher) Emit(event interface{}) {
	w.events <- event
}

// SetBufferDuration set the event buffer damping duration.
func (w *Watcher) SetBufferDuration(d time.Duration) {
	// Set the buffer duration.
	w.bufferDuration = d
}

// sendEvent send one or more events to the observer listeners.
func (w *Watcher) sendEvent(event interface{}) {
	// NOTE: Why? Why not lock it here?

	// NOTE: we do not lock this function directly.
	//
	// All functions using sendEvent must be locked
	// for operations using o.listeners.
	for _, listener := range w.listeners {
		go listener(event)
	}
}

// handleEvent handle an event.
func (w *Watcher) handleEvent(event interface{}, f *string) {
	// Lock:
	// 1. operations on listeners array (sendEvent).
	// 2. operations on bufferEvents array.
	// 3. operations using the watchPatterns set (matchFile).
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// If we do not buffer events, just send this event now.
	if w.bufferDuration == 0 {
		w.sendEvent(event)
		return
	}

	// Add new event to the event buffer.
	w.bufferEvents = append(w.bufferEvents, event)

	// If this is the first event, set a timeout function.
	if len(w.bufferEvents) == 1 {
		time.AfterFunc(w.bufferDuration, func() {
			// Lock:
			// 1. operations on listeners array (sendEvent).
			// 2. operations on bufferEvents array.
			w.mutex.Lock()
			defer w.mutex.Unlock()

			// Send all events in event buffer.
			w.sendEvent(w.bufferEvents)

			// Reset events buffer.
			w.bufferEvents = make([]interface{}, 0)
		})
	}
}

// eventLoop runs the event loop.
func (w *Watcher) eventLoop() error {
	// Run observer.
	go func() {
		for {
			select {
			case event := <-w.events:
				w.handleEvent(event, nil)
			case <-w.quit:
				return
			}
		}
	}()

	return nil
}
