package core

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// MassFile is the mass file model
type MassFile struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Date string `json:"date"`
	URL  string `json:"url"`
}

// GetID returns the ID field
func (f *MassFile) GetID() string {
	return f.ID
}

// GenerateID generates a new mass file ID
func (f *MassFile) GenerateID() {
	f.ID = uuid.New().String()
}

// Pack packs the mass file object into message
func (f *MassFile) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(f)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling role while packing")
	}

	return NewMessage(
		f.ID,
		intent,
		ModelType,
		"mass_file",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the mass file object
func (f *MassFile) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, f)
}
