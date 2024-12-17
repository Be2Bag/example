package main

import (
	"log"
	"os"

	"github.com/Be2Bag/example/config"
	registerHandler "github.com/Be2Bag/example/module/register/handler"
	registerPorts "github.com/Be2Bag/example/module/register/ports"
	registerRepository "github.com/Be2Bag/example/module/register/repository"
	registerServices "github.com/Be2Bag/example/module/register/services"
	sessionHandler "github.com/Be2Bag/example/module/session/handler"
	sessionPorts "github.com/Be2Bag/example/module/session/ports"
	sessionRepository "github.com/Be2Bag/example/module/session/repository"
	sessionServices "github.com/Be2Bag/example/module/session/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func setupEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func setupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New())
}

func main() {
	setupEnvironment()
	config.InitDatabase()

	app := fiber.New()
	setupMiddleware(app)

	apiGroup := app.Group("/api")

	v := validator.New()

	sharedRepo := registerRepository.NewRegisterRepository(config.DB, nil)

	var registerRepo registerPorts.RegisterRepository = registerRepository.NewRegisterRepository(config.DB, sharedRepo)
	var registerService registerPorts.RegisterService = registerServices.NewRegisterService(registerRepo)
	registerHandler := registerHandler.NewRegisterHandler(registerService, v)

	var sessionRepo sessionPorts.SessionRepository = sessionRepository.NewSessionRepository(config.DB, sharedRepo)
	var sessionService sessionPorts.SessionService = sessionServices.NewSessionService(sessionRepo)
	sessionHandler := sessionHandler.NewSessionHandler(sessionService, v)

	registerHandler.SetupRoutes(apiGroup)
	sessionHandler.SetupRoutes(apiGroup)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server is running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
