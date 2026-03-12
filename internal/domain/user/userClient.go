package user

type UserClient interface {
	GetUser(userId string) (*User, error)
}
