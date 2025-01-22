package handlers

import (
	"io"
	"net/http"
)

// Ping godoc
// @Summary Проверка доступности сервиса
// @Description Возвращает строку "OK" для проверки работоспособности API
// @Tags Health Check
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "OK")
}
