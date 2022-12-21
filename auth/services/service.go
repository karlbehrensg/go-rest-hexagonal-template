package auth

import "github.com/gofiber/fiber/v2"

type AuthService interface {
	CreateUser(c *fiber.Ctx) error
}
