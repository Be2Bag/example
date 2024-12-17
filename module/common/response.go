package common

import (
	"time"

	formatter "github.com/Be2Bag/example/pkg/formatter"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status    string      `json:"status"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

func NewResponse(status string, code int, message string, data interface{}) *Response {
	return &Response{
		Status:    status,
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: formatter.FormatThaiDate(time.Now().UTC()),
	}
}

func SendSuccessResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	response := NewResponse("success", code, message, data)
	return c.Status(code).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	response := NewResponse("error", code, message, data)
	return c.Status(code).JSON(response)
}
