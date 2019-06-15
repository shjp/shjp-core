package model

import (
	"github.com/shjp/shjp-core"
)

// User is the User model
type User struct {
	core.User

	Groups []*userGroup `json:"groups"`
}

type userGroup struct {
	core.User

	Status      string               `json:"status"`
	RoleName    string               `json:"role_name"`
	Privilege   int                  `json:"privilege"`
	Permissions core.GroupPermission `json:"permissions"`
}

// PopulatePermissions populates user's each group permissions from privileges
func (u *User) PopulatePermissions() error {
	for i := range u.Groups {
		u.Groups[i].Permissions = core.MapGroupPermission(u.Groups[i].Privilege)
	}

	return nil
}
