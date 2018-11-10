package model

import (
	"github.com/shjp/shjp-core"
)

// Group is the Group model
type Group struct {
	core.Group

	Members []*member `json:"members"`
}

type member struct {
	core.User

	Status    string `json:"status"`
	RoleName  string `json:"role_name"`
	Privilege int    `json:"privilege"`
}
