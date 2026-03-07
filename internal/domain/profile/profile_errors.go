package profile

import "errors"

var (
	ErrProfileNotFound = errors.New("PROFILE_NOT_FOUND")
	ErrLongDisplayName = errors.New("LONG_DISPLAY_NAME")
	ErrLongBio         = errors.New("LONG_BIO")
	ErrNotActivated    = errors.New("INACTIVE_USER")
)
