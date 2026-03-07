package profile

import "time"

const MaxBioLength int = 500
const MaxDisplayNameLength int = 32

type ProfileRepository interface {
	Create(profile *Profile) error
	Update(profile *Profile) error
	Delete(profileId string) error
	DeleteByUserId(userId string) error
	GetById(profileId string) (*Profile, error)
	GetByUserID(userID string) (*Profile, error)
	GetByUsername(username string) (*Profile, error)
}

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

func (p *Profile) ChangeBio(bio string) error {
	if len(bio) > MaxBioLength {
		return ErrLongBio
	}
	p.Bio = bio
	return nil
}

func (p *Profile) ChangeDisplayName(displayName string) error {
	if len(displayName) > MaxDisplayNameLength {
		return ErrLongDisplayName
	}
	p.DisplayName = displayName
	return nil
}

func (p *Profile) ChangeAvatar(newAvatar string) error {
	if !p.IsActive {
		return ErrNotActivated
	}
	p.AvatarUrl = newAvatar
	return nil
}
