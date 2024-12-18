package middleware

import (
	"github.com/Be2Bag/example/module/common"
	util "github.com/Be2Bag/example/pkg/crypto"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return common.SendErrorResponse(c, fiber.StatusUnauthorized, "ไม่พบ token", nil)
	}

	cryptoService := util.NewCryptoService()
	data, err := cryptoService.ValidateJWTToken(token)
	if err != nil {
		return common.SendErrorResponse(c, fiber.StatusUnauthorized, "token ไม่ถูกต้อง", nil)
	}

	c.Locals("auth", data)

	return c.Next()
}
