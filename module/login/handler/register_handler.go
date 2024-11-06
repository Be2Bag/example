// module/register/handler/register_handler.go
package handler

import (
	"github.com/Be2Bag/example/module/register/dto"
	"github.com/Be2Bag/example/module/register/ports"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
    service ports.RegisterService
}

func NewRegisterHandler(service ports.RegisterService) *RegisterHandler {
    return &RegisterHandler{
        service: service,
    }
}

func (h *RegisterHandler) Register(c *fiber.Ctx) error {
    var req dto.RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "cannot parse JSON",
        })
    }

    // Validate request
    if err := c.Validate(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    resp, err := h.service.Register(req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(resp)
}
