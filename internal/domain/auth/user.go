package auth

import "time"

type AuthUser struct {
	ID        string
	Email     string
	Username  string
	IsActive  bool
	CreatedAt time.Time
}
