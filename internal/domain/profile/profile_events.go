package profile

import "time"

type ProfileAvatarChangedEvent struct {
	NewAvatar string    `json:"new_avatar"`
	ChangedAt time.Time `json:"changed_at"`
}

type FileOrphanedEvent struct {
	FilePath   string    `json:"file_path"`
	OrphanedAt time.Time `json:"orphaned_at"`
}
