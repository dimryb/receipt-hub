package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // подключаем Swagger
	"log"
	_ "receipt-loader/docs"
	"receipt-loader/internal/db"
	"receipt-loader/internal/handlers"
)

// @title Receipt Hub API
// @version 1.0
// @description This is a sample server for Receipt Hub.
// @host localhost:8080
// @BasePath /
func main() {
	// Подключение к БД
	dsn := "postgres://postgres@localhost:5432/receipts"
	database := db.Connect(dsn)

	// Инициализация роутера
	router := gin.Default()

	// Включаем Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Регистрация маршрутов
	handlers.SetupRoutes(router, database)

	// Запуск сервера
	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
