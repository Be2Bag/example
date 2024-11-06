// config/database.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Be2Bag/example/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto migrate User model
    err = DB.AutoMigrate(&model.User{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Println("Database connection established and migrated.")
}
