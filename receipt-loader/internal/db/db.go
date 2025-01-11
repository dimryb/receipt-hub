package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Connect устанавливает соединение с базой данных и возвращает экземпляр *gorm.DB
func Connect(dsn string) *gorm.DB {
	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully")
	return db
}
