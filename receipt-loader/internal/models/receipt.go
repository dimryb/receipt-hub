package models

import "time"

// Receipt представляет данные о чеке
type Receipt struct {
	Id             uint      `json:"id" swaggerignore:"true"`
	Date           string    `json:"date" binding:"required" example:"2024-12-21"`
	Time           string    `json:"time" binding:"required" example:"14:20:00"`
	Amount         float64   `json:"amount" binding:"required" example:"1.0"`
	FiscalNumber   uint64    `json:"fiscal_number" binding:"required" example:"10"`
	FiscalDocument uint      `json:"fiscal_document" binding:"required" example:"100"`
	FiscalSign     uint64    `json:"fiscal_sign" binding:"required" example:"1"`
	CreatedAt      time.Time `json:"created_at,omitempty" swaggerignore:"true"`
}

// ErrorResponse структура для ошибок
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

// SuccessResponse структура для успешных ответов
type SuccessResponse struct {
	Message string `json:"message"`
	ID      uint   `json:"id"`
}
