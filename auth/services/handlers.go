package auth

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/karlbehrensg/go-rest-template/auth/domain"
)

type authService struct {
	authRepo *domain.AuthRepository
}

func NewAuthService(authRepo *domain.AuthRepository) AuthService {
	service := &authService{
		authRepo: authRepo,
	}
	return service
}

func (service *authService) CreateUser(c *fiber.Ctx) error {
	data := &CreateUser{}
	if err := c.BodyParser(data); err != nil {
		return err
	}

	if data.Password != data.Password2 {
		return fiber.NewError(fiber.StatusBadRequest, "Passwords do not match")
	}

	user := &domain.User{
		Email:    data.Email,
		Password: data.Password,
	}

	err := (*service.authRepo).CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User created",
	})
}
