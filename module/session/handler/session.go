package handler

import (
	"errors"
	"time"

	"github.com/Be2Bag/example/middleware"
	"github.com/Be2Bag/example/module/common"
	"github.com/Be2Bag/example/module/session/dto"
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
	protected.Get("/", h.Session)
}

func (h *SessionHandler) Login(c *fiber.Ctx) error {
	var sessionRequest dto.SessionRequest
	if err := c.BodyParser(&sessionRequest); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	token, err := h.sessionService.Login(sessionRequest)
	if err != nil {
		if errors.Is(err, services.ErrInvalidEmail) {
			return common.SendErrorResponse(c, fiber.StatusUnauthorized, "อีเมลไม่ถูกต้อง", nil)
		}
		if errors.Is(err, services.ErrInvalidPassword) {
			return common.SendErrorResponse(c, fiber.StatusUnauthorized, "รหัสผ่านไม่ถูกต้อง", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true
	cookie.SameSite = "Strict"
	c.Cookie(cookie)

	return common.SendSuccessResponse(c, fiber.StatusOK, "เข้าสู่ระบบสำเร็จ", nil)
}

func (h *SessionHandler) Session(c *fiber.Ctx) error {
	auth := c.Locals("auth").(map[string]interface{})
	return common.SendSuccessResponse(c, fiber.StatusOK, "ตรวจสอบเซสชันสำเร็จ", auth)
}
