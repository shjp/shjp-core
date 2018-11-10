package core

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// User is the user data model
type User struct {
	ID            string  `json:"id"`
	Name          *string `json:"name"`
	BaptismalName *string `json:"baptismal_name"`
	Birthday      *string `json:"birthday"`
	FeastDay      *string `json:"feast_day"`
	Created       *string `json:"created"`
	LastActive    *string `json:"last_active"`
}

// GetID returns the ID field
func (u *User) GetID() string {
	return u.ID
}

// GenerateID generates a new user ID
func (u *User) GenerateID() {
	u.ID = uuid.New().String()
}

// Pack packs the user object into message
func (u *User) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(u)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling user while packing")
	}

	return NewMessage(
		u.ID,
		intent,
		ModelType,
		"user",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the user object
func (u *User) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, u)
}

/*func (u *User) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}*/
