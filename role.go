package core

// Role is the role model (no pun intended..)
type Role struct {
	ID string `json:"id"`
	// Because of a super weird go-pg behaviour that doesn't seem to parse
	// this struct when we include GroupID, omit it for now
	//Group_id  string `json:"group_id"`
	Name      string `json:"name"`
	Privilege int    `json:"privilege"`
}
