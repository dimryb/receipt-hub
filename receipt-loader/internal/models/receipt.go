package models

import "time"

// Receipt представляет данные о чеке
type Receipt struct {
	ID             int       `json:"id"`
	Date           string    `json:"date" binding:"required"` // Формат: YYYY-MM-DD
	Time           string    `json:"time" binding:"required"` // Формат: HH:mm
	Amount         float64   `json:"amount" binding:"required"`
	FiscalNumber   int64     `json:"fiscal_number" binding:"required"`
	FiscalDocument int       `json:"fiscal_document" binding:"required"`
	FiscalSign     int64     `json:"fiscal_sign" binding:"required"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
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
