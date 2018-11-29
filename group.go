package core

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Group is the group data model
type Group struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// GetID returns the ID field
func (g *Group) GetID() string {
	return g.ID
}

// GenerateID generates a new group ID
func (g *Group) GenerateID() {
	g.ID = uuid.New().String()
}

// Pack packs the group object into message
func (g *Group) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(g)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling group while packing")
	}

	return NewMessage(
		g.ID,
		intent,
		ModelType,
		"group",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the group object
func (g *Group) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, g)
}
