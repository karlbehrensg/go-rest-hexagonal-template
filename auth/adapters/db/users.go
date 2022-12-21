package auth

import (
	"database/sql"

	domain "github.com/karlbehrensg/go-rest-template/auth/domain"
)

type userAdapter struct {
	client *sql.DB
}

func NewUserAdapter(client *sql.DB) (domain.AuthRepository, error) {
	adapter := &userAdapter{}

	adapter.client = client

	return adapter, nil
}

func (adapter *userAdapter) CreateUser(user *domain.User) error {
	sqlStatement := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)`
	_, err := adapter.client.Exec(sqlStatement, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
