package core

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Role is the role model (no pun intended..)
type Role struct {
	ID string `json:"id"`
	// Because of a super weird go-pg behaviour that doesn't seem to parse
	// this struct when we include GroupID, omit it for now
	GroupID   string `json:"group_id"`
	Name      string `json:"name"`
	Privilege int    `json:"privilege"`
}

// GetID returns the ID field
func (r *Role) GetID() string {
	return r.ID
}

// GenerateID generates a new group ID
func (r *Role) GenerateID() {
	r.ID = uuid.New().String()
}

// Pack packs the group object into message
func (r *Role) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(r)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling role while packing")
	}

	return NewMessage(
		r.ID,
		intent,
		ModelType,
		"role",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the group object
func (r *Role) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, r)
}
