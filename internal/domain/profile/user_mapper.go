package profile

import "main/internal/domain/auth"

func MapUserToProfile(user *auth.AuthUser) *Profile {
	return &Profile{
		UserId:      user.ID,
		Username:    user.Username,
		DisplayName: user.Username,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt,
	}
}
