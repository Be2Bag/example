package handler

import (
	"errors"
	"time"

	"github.com/Be2Bag/example/module/session/dto"
	"github.com/Be2Bag/example/module/session/middleware"
	"github.com/Be2Bag/example/module/session/ports"
	"github.com/Be2Bag/example/module/session/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SessionHandler struct {
	sessionService ports.SessionService
	validator      *validator.Validate
}

func NewSessionHandler(sessionService ports.SessionService, v *validator.Validate) *SessionHandler {

	return &SessionHandler{
		sessionService: sessionService,
		validator:      v,
	}
}

func (h *SessionHandler) SetupRoutes(router fiber.Router) {
	session := router.Group("/session")
	session.Post("/login", h.Login)

	protected := router.Group("/session", middleware.AuthMiddleware)
	protected.Get("/", h.Test)
}

func (h *SessionHandler) Login(c *fiber.Ctx) error {
	var sessionRequest dto.SessionRequest
	if err := c.BodyParser(&sessionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := h.sessionService.Login(sessionRequest)
	if err != nil {
		if errors.Is(err, services.ErrInvalidEmail) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "อีเมล์ไม่ถูกต้อง",
			})
		}
		if errors.Is(err, services.ErrInvalidPassword) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "รหัสผ่านไม่ถูกต้อง",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์",
			"details": err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}

func (h *SessionHandler) Test(c *fiber.Ctx) error {
	auth := c.Locals("auth").(map[string]interface{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Session successful",
		"user":    auth,
	})
}
