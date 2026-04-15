package database

import (
	"learning-backend/models"
	"log"
"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	fmt.Println("Database connected successfully")

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Cart{},
	)
}