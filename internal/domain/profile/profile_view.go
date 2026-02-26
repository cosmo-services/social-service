package profile

import "time"

type ProfileView struct {
	ID          string    `json:"id"`
	UserId      string    `json:"user_id"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	AvatarUrl   string    `json:"avatar_url,omitempty"`
	IsActive    bool      `json:"is_active,omitempty"`
	IsDeleted   bool      `json:"is_deleted,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
