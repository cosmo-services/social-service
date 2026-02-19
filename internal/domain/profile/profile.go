package profile

type Profile struct {
	ID          string `json:"id"`
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"display_name"`
	AvatarUrl   string `json:"avatar_url"`
}

type ProfileRepository interface {
	Create(profile *Profile) error
	Update(profile *Profile) error
	Delete(profileId string) error
	GetById(profileId string) error
}
