package model

import (
	"github.com/shjp/shjp-core"
)

// Announcement is the announcement model
type Announcement struct {
	core.Announcement

	Creator core.User `json:"creator"`
}
