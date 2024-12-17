package middleware

import (
	util "github.com/Be2Bag/example/pkg/crypto"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "ไม่พบ token",
		})
	}

	data, err := util.ValidateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token ไม่ถูกต้อง",
		})
	}

	c.Locals("auth", data)

	return c.Next()
}
