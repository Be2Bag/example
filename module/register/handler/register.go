package handler

import (
	"errors"

	"github.com/Be2Bag/example/module/register/dto"
	"github.com/Be2Bag/example/module/register/ports"
	"github.com/Be2Bag/example/module/register/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// RegisterHandler คือฮันเดิลสำหรับการลงทะเบียนผู้ใช้
type RegisterHandler struct {
	registerService ports.RegisterService
	validator       *validator.Validate
}

// NewRegisterHandler สร้าง RegisterHandler ใหม่
func NewRegisterHandler(registerService ports.RegisterService, v *validator.Validate) *RegisterHandler {
	return &RegisterHandler{
		registerService: registerService,
		validator:       v,
	}
}

// Register จัดการคำขอลงทะเบียนผู้ใช้ใหม่
func (h *RegisterHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ข้อมูลคำขอไม่ถูกต้อง",
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "การตรวจสอบล้มเหลว",
			"details": err.Error(),
		})
	}

	resp, err := h.registerService.Register(req)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "ผู้ใช้มีอยู่แล้ว",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}
