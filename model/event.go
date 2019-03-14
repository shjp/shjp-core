package model

import core "github.com/shjp/shjp-core"

// Event is the event data model
type Event struct {
	core.Event

	Author core.User `json:"author"`
	Rsvps  []*rsvp   `json:"rsvps"`
}

type rsvp struct {
	User core.User `json:"user"`
	RSVP string    `json:"rsvp"`
}
