package auth

type AuthRepository interface {
	CreateUser(user *User) error
}
