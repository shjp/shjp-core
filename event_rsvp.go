package core

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// EventRSVP represents the user's RSVP state to the event
type EventRSVP struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	RSVP    string `json:"rsvp"`
}

// Pack packs the event RSVP object into message
func (r *EventRSVP) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling event RSVP while packing")
	}

	return NewMessage(
		r.EventID+r.UserID,
		intent,
		RelationshipType,
		"event_rsvp",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the event RSVP object
func (r *EventRSVP) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, r)
}
