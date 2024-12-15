package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDatabase เชื่อมต่อกับฐานข้อมูล MongoDB
func ConnectDatabase() (*mongo.Database, error) {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	uri := os.Getenv("DB_URI") // รับ URI ของฐานข้อมูลจาก environment
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	database := client.Database(os.Getenv("DB_NAME")) // เลือกฐานข้อมูล
	return database, nil
}

// InitDatabase เริ่มต้นการเชื่อมต่อฐานข้อมูล
func InitDatabase() {
	var err error
	DB, err = ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection to MongoDB established.")
}
