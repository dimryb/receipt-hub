package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"receipt-loader/internal/db"
	"receipt-loader/internal/handlers"
)

func main() {
	// Подключение к БД
	dsn := "postgres://postgres@localhost:5432/receipts"
	database := db.Connect(dsn)

	// Инициализация роутера
	router := gin.Default()

	// Регистрация маршрутов
	handlers.SetupRoutes(router, database)

	// Запуск сервера
	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
