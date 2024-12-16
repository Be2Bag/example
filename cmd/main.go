// cmd/main.go
package main

import (
	"log"
	"os"

	"github.com/Be2Bag/example/config"
	"github.com/Be2Bag/example/module/register/handler"
	"github.com/Be2Bag/example/module/register/ports"
	"github.com/Be2Bag/example/module/register/repository"
	"github.com/Be2Bag/example/module/register/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// setupEnvironment ตั้งค่าการโหลด environment variables
func setupEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// setupMiddleware ตั้งค่า middleware สำหรับแอป Fiber
func setupMiddleware(app *fiber.App) {
	app.Use(recover.New()) // ใช้ middleware สำหรับจัดการข้อผิดพลาด
	app.Use(logger.New())  // ใช้ middleware สำหรับล็อกคำขอ
}

// setupRoutes ตั้งค่าเส้นทางสำหรับแอป Fiber
func setupRoutes(app *fiber.App, registerHandler *handler.RegisterHandler) {
	api := app.Group("/api")
	api.Get("/register", registerHandler.GetUser)           // เส้นทางสำหรับแสดงผู้ใช้
	api.Get("/register/:id", registerHandler.GetUserByID)   // เส้นทางสำหรับแสดงผู้ใช้โดยใช้รหัสผู้ใช้
	api.Post("/register", registerHandler.Register)         // เส้นทางสำหรับลงทะเบียนผู้ใช้
	api.Put("/register", registerHandler.UpdateUser)        // เส้นทางสำหรับอัปเดตข้อมูลผู้ใช้
	api.Delete("/register/:id", registerHandler.DeleteUser) // เส้นทางสำหรับลบผู้ใช้
}

func main() {
	setupEnvironment()    // ตั้งค่าการโหลด environment
	config.InitDatabase() // เริ่มต้นการเชื่อมต่อฐานข้อมูล

	app := fiber.New()
	setupMiddleware(app) // ตั้งค่า middleware

	v := validator.New() // สร้าง validator ใหม่

	// ตั้งค่า repository และ service สำหรับการลงทะเบียน
	var registerRepo ports.RegisterRepository = repository.NewRegisterRepository(config.DB)
	var registerService ports.RegisterService = services.NewRegisterService(registerRepo)
	registerHandler := handler.NewRegisterHandler(registerService, v)

	setupRoutes(app, registerHandler) // ตั้งค่าเส้นทาง

	// กำหนดพอร์ตที่เซิร์ฟเวอร์จะรัน
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server is running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
