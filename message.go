package core

import (
	"fmt"
)

// Message is an encapsulated unit for the queue communication
type Message struct {
	Key           string        `json:"key"`
	Intent        Intent        `json:"intent"`
	Type          MessageType   `json:"type"`
	Subtype       string        `json:"subtype"`
	OperationType OperationType `json:"operation_type"`
	Data          []byte        `json:"data"`
	Error         error         `json:"error"`
}

// NewMessage creates a new message object
func NewMessage(id string, i Intent, mt MessageType, st string, op OperationType, blob []byte, err error) *Message {
	m := &Message{
		Intent:        i,
		Type:          mt,
		Subtype:       st,
		OperationType: op,
		Data:          blob,
		Error:         err}
	m.generateKey(id)
	return m
}

// ReadFailureMessage creates a message whose content could not be read
func ReadFailureMessage(blob []byte, err error) *Message {
	return GenericFailureMessage("", UnknownType, "", UnknownOperation, blob, err)
}

// TypedFailureMessage creates a message whose type is known
func TypedFailureMessage(mt MessageType, blob []byte, err error) *Message {
	return GenericFailureMessage("", mt, "", UnknownOperation, blob, err)
}

// GenericFailureMessage creates a message that indicates a generic failure
func GenericFailureMessage(id string, mt MessageType, st string, op OperationType, blob []byte, err error) *Message {
	return NewMessage(id, IntentFailure, mt, st, op, blob, err)
}

// ExtractEntity extracts the entity object from the message
func (m *Message) ExtractEntity() (Entity, error) {
	var entity Entity
	switch m.Subtype {
	case "announcement":
		entity = &Announcement{}
	case "event":
		entity = &Event{}
	case "event_rsvp":
		entity = &EventRSVP{}
	case "group":
		entity = &Group{}
	case "group_membership":
		entity = &GroupMembership{}
	case "role":
		entity = &Role{}
	case "user":
		entity = &User{}
	default:
		return nil, fmt.Errorf("Trying to extract entity with unrecognized subtype %s", m.Subtype)
	}
	err := entity.Unpack(*m)
	return entity, err
}

func (m *Message) generateKey(id string) {
	m.Key = fmt.Sprintf("%s.%s.%s.%s.%s", m.Intent, m.Type, m.Subtype, m.OperationType, id)
}
