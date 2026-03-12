package profile

import user_domain "main/internal/domain/user"

func MapUserToProfile(user *user_domain.User) *Profile {
	return &Profile{
		UserId:      user.ID,
		Username:    user.Username,
		DisplayName: user.Username,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt,
	}
}
