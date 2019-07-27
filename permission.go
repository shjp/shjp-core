package core

// Permission enums
const (
	PermissionRead = 1 << iota
	PermissionMembersRead
	PermissionCommentRead
	PermissionCommentModerate
	PermissionAnnouncementModerate
	PermissionEventModerate
	PermissionAdmin
	PermissionSuper
)

// GroupPermission defines user's permissions within a group
type GroupPermission struct {
	CanRead               bool `json:"can_read"`
	CanReadMembers        bool `json:"can_read_members"`
	CanReadComments       bool `json:"can_read_comments"`
	CanWriteComments      bool `json:"can_write_comments"`
	CanWriteAnnouncements bool `json:"can_write_announcements"`
	CanWriteEvents        bool `json:"can_write_events"`
	CanAdminGroup         bool `json:"can_admin_group"`
	CanEditUsers          bool `json:"can_edit_users"`
}

// MapGroupPermission maps out the permissions from the privilege integer
func MapGroupPermission(priv int) GroupPermission {
	return GroupPermission{
		CanRead:               priv&PermissionRead != 0,
		CanReadMembers:        priv&PermissionMembersRead != 0,
		CanReadComments:       priv&PermissionCommentRead != 0,
		CanWriteComments:      priv&PermissionCommentModerate != 0,
		CanWriteAnnouncements: priv&PermissionAnnouncementModerate != 0,
		CanWriteEvents:        priv&PermissionEventModerate != 0,
		CanAdminGroup:         priv&PermissionAdmin != 0,
		CanEditUsers:          priv&PermissionSuper != 0,
	}
}
