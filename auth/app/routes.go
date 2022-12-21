package auth

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/karlbehrensg/go-rest-template/auth/domain"
	service "github.com/karlbehrensg/go-rest-template/auth/services"
)

func ApplyRoutes(app *fiber.App, authRepo *domain.AuthRepository) {
	authService := service.NewAuthService(authRepo)

	authRoutes := app.Group("/auth")
	authRoutes.Post("/signup", authService.CreateUser)
}
