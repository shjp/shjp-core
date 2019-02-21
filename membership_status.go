package core

// MembershipStatus is enum for group membership status
type MembershipStatus string

const (
	// Accepted represents accepted status
	Accepted MembershipStatus = "accepted"

	// Pending represents pending status
	Pending MembershipStatus = "pending"
)
