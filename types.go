package watch

import (
	"io"
)

// Node represents a struct that has a unique
// id, and a status property returning a constant
type Node struct {
	id int `json:"id,omitempty"`

	status string `json:"status,omitempty"`
}

// Device is a representation of a watch
type Device struct {
	id int `json:"id,omitempty"`

	status   string `json:"status,omitempty"`
	metadata []byte `json:"data,omitempty"` // the devices hardware info
}

// Event mirrors the Node struct but also
// holds `n` bytes of data associated with it
type Event struct {
	id int `json:"id,omitempty"`

	status string `json:"status,omitempty"`
	data   []byte `json:"data,omitempty"`
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

// Page is a less urgent WatcherEmitter
// instead of Alerts, Page implements Notify()
type Pager interface {
	Notify(Event) error
	Watcher
}
