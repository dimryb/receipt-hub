package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddReceipt сохраняет данные о чеке
func AddReceipt(c *gin.Context, db *gorm.DB) {

}

// GetReceiptByID возвращает чек по его ID
func GetReceiptByID(c *gin.Context, db *gorm.DB) {

}

// SetupRoutes инициализирует маршруты API
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/receipts", func(c *gin.Context) {
		AddReceipt(c, db)
	})
	router.GET("/receipts/:id", func(c *gin.Context) {
		GetReceiptByID(c, db)
	})
}
