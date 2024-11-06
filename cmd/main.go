// cmd/main.go
package main

import (
	"log"
	"os"

	"github.com/Be2Bag/example/config"
	"github.com/Be2Bag/example/module/register/handler"
	"github.com/Be2Bag/example/module/register/repository"
	"github.com/Be2Bag/example/module/register/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/go-playground/validator/v10"
)

// CustomValidator implements fiber's Validator interface
type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // Initialize database
    config.InitDatabase()

    // Initialize Fiber app with custom validator
    app := fiber.New(fiber.Config{
        Validator: &CustomValidator{validator: validator.New()},
    })

    // Middleware
    app.Use(recover.New())
    app.Use(logger.New())

    // Setup Register module
    registerRepo := repository.NewRegisterRepository(config.DB)
    registerService := services.NewRegisterService(registerRepo)
    registerHandler := handler.NewRegisterHandler(registerService)

    // Routes
    api := app.Group("/api")
    api.Post("/register", registerHandler.Register)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    log.Printf("Server is running on port %s", port)
    log.Fatal(app.Listen(":" + port))
}
