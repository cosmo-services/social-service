package user

import "time"

type User struct {
	ID        string
	Email     string
	Username  string
	IsActive  bool
	CreatedAt time.Time
}
