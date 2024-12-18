package handler

import (
	"errors"

	"github.com/Be2Bag/example/module/common"
	"github.com/Be2Bag/example/module/register/dto"
	"github.com/Be2Bag/example/module/register/ports"
	"github.com/Be2Bag/example/module/register/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	registerService ports.RegisterService
	validator       *validator.Validate
}

func NewRegisterHandler(registerService ports.RegisterService, v *validator.Validate) *RegisterHandler {
	return &RegisterHandler{
		registerService: registerService,
		validator:       v,
	}
}

func (h *RegisterHandler) SetupRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Post("/register", h.Register)
	users.Get("/", h.GetUser)
	users.Get("/:id", h.GetUserByID)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}

func (h *RegisterHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "ข้อมูลคำขอไม่ถูกต้อง", nil)
	}

	if err := h.validator.Struct(req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "การตรวจสอบล้มเหลว", err.Error())
	}

	resp, err := h.registerService.Register(req)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			return common.SendErrorResponse(c, fiber.StatusConflict, "ผู้ใช้มีอยู่แล้ว", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", nil)
	}

	return common.SendSuccessResponse(c, fiber.StatusCreated, "ลงทะเบียนสำเร็จ", resp)
}

func (h *RegisterHandler) GetUser(c *fiber.Ctx) error {
	resp, err := h.registerService.GetUsers()
	if err != nil {
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", nil)
	}

	return common.SendSuccessResponse(c, fiber.StatusOK, "ดึงข้อมูลผู้ใช้สำเร็จ", resp)
}

func (h *RegisterHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	resp, err := h.registerService.GetUserByID(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return common.SendErrorResponse(c, fiber.StatusNotFound, "ไม่พบผู้ใช้ตาม ID ที่ระบุ", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	return common.SendSuccessResponse(c, fiber.StatusOK, "ดึงข้อมูลผู้ใช้สำเร็จ", resp)
}

func (h *RegisterHandler) UpdateUser(c *fiber.Ctx) error {
	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "ข้อมูลคำขอไม่ถูกต้อง", nil)
	}

	if err := h.validator.Struct(req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "การตรวจสอบล้มเหลว", err.Error())
	}

	resp, err := h.registerService.UpdateUser(req)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return common.SendErrorResponse(c, fiber.StatusNotFound, "ไม่พบผู้ใช้ตาม ID ที่ระบุ", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	return common.SendSuccessResponse(c, fiber.StatusOK, "อัปเดตผู้ใช้สำเร็จ", resp)
}

func (h *RegisterHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.registerService.DeleteUser(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return common.SendErrorResponse(c, fiber.StatusNotFound, "ไม่พบผู้ใช้ตาม ID ที่ระบุ", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	return common.SendSuccessResponse(c, fiber.StatusNoContent, "ลบผู้ใช้สำเร็จ", nil)
}
