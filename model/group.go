package model

import (
	"github.com/shjp/shjp-core"
)

// Group is the Group model
type Group struct {
	core.Group

	Members []*member    `json:"members"`
	Roles   []*groupRole `json:"roles"`
}

// PopulateRolesPermissions populates each of group role's permissions from their privileges
func (g *Group) PopulateRolesPermissions() {
	for i := range g.Roles {
		g.Roles[i].PopulatePermissions()
	}
}

type member struct {
	core.User

	Status    string `json:"status"`
	RoleName  string `json:"role_name"`
	Privilege int    `json:"privilege"`
}

type groupRole struct {
	core.Role

	Permissions core.GroupPermission `json:"permissions"`
}

func (gr *groupRole) PopulatePermissions() {
	gr.Permissions = core.MapGroupPermission(gr.Role.Privilege)
}
