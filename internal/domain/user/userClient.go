package user

type UserClient interface {
	GetUserById(userId string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}
