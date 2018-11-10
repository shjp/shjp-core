package core

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Announcement is the announcement data model
type Announcement struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	AuthorID string `json:"authorId"`
	Content  string `json:"content"`
}

// GetID returns the ID field
func (a *Announcement) GetID() string {
	return a.ID
}

// GenerateID generates a new announcement ID
func (a *Announcement) GenerateID() {
	a.ID = uuid.New().String()
}

// Pack packs the announcement object into message
func (a *Announcement) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling announcement while packing")
	}

	return NewMessage(
		a.ID,
		intent,
		ModelType,
		"announcement",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the announcement object
func (a *Announcement) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, a)
}

/*func (a *Announcement) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, a)
}*/
