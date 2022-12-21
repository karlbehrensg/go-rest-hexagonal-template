package auth

type CreateUser struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Password2 string `json:"password2" validate:"required"`
}
