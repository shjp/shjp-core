package core

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// GroupMembership represents user's membership within group
type GroupMembership struct {
	GroupID string           `json:"group_id"`
	UserID  string           `json:"user_id"`
	RoleID  string           `json:"role_id"`
	Status  MembershipStatus `json:"status"`
}

// Pack packs the group membership object into message
func (m *GroupMembership) Pack(intent Intent, operation OperationType) (*Message, error) {
	blob, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling group membership while packing")
	}

	return NewMessage(
		m.GroupID+m.UserID,
		intent,
		RelationshipType,
		"group_membership",
		operation,
		blob,
		nil), nil
}

// Unpack unpacks the message into the group membership object
func (m *GroupMembership) Unpack(msg Message) error {
	return json.Unmarshal(msg.Data, m)
}
