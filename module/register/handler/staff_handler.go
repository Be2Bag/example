package handler

import (
	"errors"
	"strings"

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
	users := router.Group("/staffs")
	users.Post("/register", h.Register)
	users.Get("/", h.GetStaff)
	users.Get("/:user_id", h.GetStaffByID)
	users.Put("/:user_id", h.UpdateStaff)
	users.Delete("/:user_id", h.DeleteStaff)
}

func (h *RegisterHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "ข้อมูลคำขอไม่ถูกต้อง", nil)
	}

	if err := h.validator.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
			case "Password":
				switch fieldError.Tag() {
				case "min":
					errorMessages = append(errorMessages, "รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร")
				case "max":
					errorMessages = append(errorMessages, "รหัสผ่านต้องมีความยาวไม่เกิน 64 ตัวอักษร")
				case "containsany":
					if strings.Contains(fieldError.Param(), "!@#$%^&*()") {
						errorMessages = append(errorMessages, "รหัสผ่านต้องมีอักขระพิเศษอย่างน้อยหนึ่งตัวจากกลุ่ม !@#$%^&*()")
					} else if strings.Contains(fieldError.Param(), "0123456789") {
						errorMessages = append(errorMessages, "รหัสผ่านต้องมีตัวเลขอย่างน้อยหนึ่งตัวจากกลุ่ม 0123456789")
					} else if strings.Contains(fieldError.Param(), "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
						errorMessages = append(errorMessages, "รหัสผ่านต้องมีตัวอักษรตัวใหญ่ (A-Z) อย่างน้อยหนึ่งตัว")
					} else if strings.Contains(fieldError.Param(), "abcdefghijklmnopqrstuvwxyz") {
						errorMessages = append(errorMessages, "รหัสผ่านต้องมีตัวอักษรตัวเล็ก (a-z) อย่างน้อยหนึ่งตัว")
					}
				}
			default:
				errorMessages = append(errorMessages, fieldError.Error())
			}
		}
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "การตรวจสอบล้มเหลว", strings.Join(errorMessages, ", "))
	}

	resp, err := h.registerService.Register(req)
	if err != nil {
		if errors.Is(err, services.ErrStaffAlreadyExists) {
			return common.SendErrorResponse(c, fiber.StatusConflict, "ผู้ใช้มีอยู่แล้ว", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", nil)
	}

	return common.SendSuccessResponse(c, fiber.StatusCreated, "ลงทะเบียนสำเร็จ", resp)
}

func (h *RegisterHandler) GetStaff(c *fiber.Ctx) error {
	resp, err := h.registerService.GetStaffs()
	if err != nil {
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", nil)
	}

	return common.SendSuccessResponse(c, fiber.StatusOK, "ดึงข้อมูลผู้ใช้สำเร็จ", resp)
}

func (h *RegisterHandler) GetStaffByID(c *fiber.Ctx) error {
	id := c.Params("user_id")
	resp, err := h.registerService.GetStaffByID(id)
	if err != nil {
		if errors.Is(err, services.ErrStaffNotFound) {
			return common.SendErrorResponse(c, fiber.StatusNotFound, "ไม่พบผู้ใช้ตาม ID ที่ระบุ", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	return common.SendSuccessResponse(c, fiber.StatusOK, "ดึงข้อมูลผู้ใช้สำเร็จ", resp)
}

func (h *RegisterHandler) UpdateStaff(c *fiber.Ctx) error {
	var req dto.UpdateStaffRequest
	if err := c.BodyParser(&req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "ข้อมูลคำขอไม่ถูกต้อง", nil)
	}

	if err := h.validator.Struct(req); err != nil {
		return common.SendErrorResponse(c, fiber.StatusBadRequest, "การตรวจสอบล้มเหลว", err.Error())
	}

	resp, err := h.registerService.UpdateStaff(req)
	if err != nil {
		if errors.Is(err, services.ErrStaffNotFound) {
			return common.SendErrorResponse(c, fiber.StatusNotFound, "ไม่พบผู้ใช้ตาม ID ที่ระบุ", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	return common.SendSuccessResponse(c, fiber.StatusOK, "อัปเดตผู้ใช้สำเร็จ", resp)
}

func (h *RegisterHandler) DeleteStaff(c *fiber.Ctx) error {
	id := c.Params("user_id")
	err := h.registerService.DeleteStaff(id)
	if err != nil {
		if errors.Is(err, services.ErrStaffNotFound) {
			return common.SendErrorResponse(c, fiber.StatusNotFound, "ไม่พบผู้ใช้ตาม ID ที่ระบุ", nil)
		}
		return common.SendErrorResponse(c, fiber.StatusInternalServerError, "เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์", err.Error())
	}

	return common.SendSuccessResponse(c, fiber.StatusNoContent, "ลบผู้ใช้สำเร็จ", nil)
}
