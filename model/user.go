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

	Status    string `json:"status"`
	RoleName  string `json:"role_name"`
	Privilege int    `json:"privilege"`
}
