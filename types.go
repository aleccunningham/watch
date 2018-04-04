package watch

import (
	"io"
)

// Node represents a struct that has a unique
// id, and a status property returning a constant
type Node struct {
	id int

	status string
}

// Device is a representation of a watch
type Device struct {
	id int

	status   string
	metadata []byte // the devices hardware info
}

// Event mirrors the Node struct but also
// holds `n` bytes of data associated with it
type Event struct {
	id int

	status string
	data   []byte
}

// BatchRequest is a slice of Events that are
// collected and sent to the Watcher together
type BatchRequest []Event

// Watcher ingests incoming events
// via event specific handlers
type Watcher interface {
	Handle(Event) error
}

// Emitter is used to send alerts via
// a BatchRequest or individual devices
type Emitter interface {
	Alert() (Event, error)
}

// WatcherEmitter can handler events, along
// with emitter its own events
type WatcherEmitter interface {
	Watcher
	Emitter
}

// Pager is a less urgent WatcherEmitter
// instead of Alerts, Page implements Notify()
type Pager interface {
	Notify(Event) error
	Watcher
}
