package core

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Event is the event model
type Event struct {
	ID                  string     `json:"id"`
	Name                string     `json:"name"`
	Start               *string    `json:"start"`
	End                 *string    `json:"end"`
	Creator             *string    `json:"creator"`
	Deadline            *string    `json:"deadline"`
	AllowMaybe          bool       `json:"allow_maybe" sql:",notnull"`
	Description         string     `json:"description"`
	Location            *string    `json:"location"`
	LocationDescription *string    `json:"location_description"`
	Created             *time.Time `json:"created"`
	Updated             *time.Time `json:"updated"`

	GroupIDs []string `json:"group_ids" sql:"-"`
}

// GetID returns the ID field
func (e *Event) GetID() string {
	return e.ID
}

// GenerateID generates a new event ID
func (e *Event) GenerateID() {
	e.ID = uuid.New().String()
}

// Pack packs the event object into message
func (e *Event) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling event while packing")
	}

	return NewMessage(
		e.ID,
		intent,
		ModelType,
		"event",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the event object
func (e *Event) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, e)
}
