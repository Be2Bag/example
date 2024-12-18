package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ResponseTimeMiddleware(c *fiber.Ctx) error {
	startTime := time.Now()
	if err := c.Next(); err != nil {
		return err
	}
	duration := time.Since(startTime)
	if duration > time.Second {
		log.Printf("⚠️   Warning: %s %s took %v ⚠️", c.Method(), c.OriginalURL(), duration)
	}
	return nil
}
