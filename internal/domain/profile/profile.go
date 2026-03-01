package profile

import "time"

type Profile struct {
	ID          string    `json:"id"`
	UserId      string    `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	AvatarUrl   string    `json:"avatar_url"`
	Bio         string    `json:"bio"`
	IsActive    bool      `json:"is_active"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProfileRepository interface {
	Create(profile *Profile) error
	Update(profile *Profile) error
	Delete(profileId string) error
	GetById(profileId string) (*Profile, error)
	GetByUserID(userID string) (*Profile, error)
	GetByUsername(username string) (*Profile, error)
}
