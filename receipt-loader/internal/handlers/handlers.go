package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"receipt-loader/internal/models"
	"time"
)

// AddReceipt сохраняет данные о чеке
// @Summary Add a new receipt
// @Description Save a new receipt in the database. If a duplicate receipt exists, returns a conflict error.
// @Tags Receipts
// @Accept json
// @Produce json
// @Param receipt body models.Receipt true "Receipt details"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse "Некорректный запрос"
// @Failure 409 {object} models.ErrorResponse "Чек уже существует"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера"
// @Router /receipts [post]
func AddReceipt(c *gin.Context, db *gorm.DB) {
	var receipt models.Receipt

	// Привязка JSON к модели
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Проверка на существование дубликата по ключевым полям
	existingReceipt := models.Receipt{}
	err := db.Where("fiscal_number = ? AND fiscal_document = ? AND fiscal_sign = ?",
		receipt.FiscalNumber,
		receipt.FiscalDocument,
		receipt.FiscalSign,
	).First(&existingReceipt).Error
	if err == nil {
		// Дубликат найден
		c.JSON(http.StatusConflict, gin.H{
			"error": "Receipt already exists",
			"id":    existingReceipt.ID,
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		// Ошибка при проверке
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to verify receipt uniqueness",
		})
		return
	}

	// Установка времени создания
	receipt.CreatedAt = time.Now()

	// Сохранение данных в базу
	if err := db.Create(&receipt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save receipt",
			"details": err.Error(),
		})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Receipt saved successfully",
		"id":      receipt.ID,
	})
}

// GetReceiptByID возвращает чек по его ID
// @Summary Get receipt by ID
// @Description Retrieve a receipt by its unique ID.
// @Tags Receipts
// @Accept json
// @Produce json
// @Param id path string true "Receipt ID"
// @Success 200 {object} models.Receipt
// @Failure 404 {object} models.ErrorResponse "Чек не найден"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера"
// @Router /receipts/{id} [get]
func GetReceiptByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var receipt models.Receipt

	// Поиск чека по ID
	if err := db.First(&receipt, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Receipt not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve receipt",
			})
		}
		return
	}

	c.JSON(http.StatusOK, receipt)
}

// SetupRoutes инициализирует маршруты API
// @Summary Initialize routes for receipts API
// @Description Setup all routes for handling receipts API requests.
// @Tags Receipts
// @Router /receipts [post]
// @Router /receipts/{id} [get]
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/receipts", func(c *gin.Context) {
		AddReceipt(c, db)
	})
	router.GET("/receipts/:id", func(c *gin.Context) {
		GetReceiptByID(c, db)
	})
}
