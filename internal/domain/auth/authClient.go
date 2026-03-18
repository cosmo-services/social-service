package auth

type AuthClient interface {
	GetUserById(userId string) (*AuthUser, error)
	GetUserByUsername(username string) (*AuthUser, error)
}
