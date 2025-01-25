package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"receipt-loader/internal/models"
	"receipt-loader/internal/rest"
	"strconv"
)

// AddReceipt сохраняет данные о чеке
// @Summary Add a new receipt
// @Description Save a new receipt in the database. If a duplicate receipt exists, returns a conflict error.
// @Tags Receipt
// @Accept json
// @Produce json
// @Param receipt body models.Receipt true "Receipt details"
// @Success 201 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse "Некорректный запрос"
// @Failure 409 {object} models.ErrorResponse "Чек уже существует"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера"
// @Router /receipt [post]
func AddReceipt(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var receipt models.Receipt

		err := json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}

		// Проверка на существование дубликата по ключевым полям
		existingReceipt := models.Receipt{}
		err = db.Where("fiscal_number = ? AND fiscal_document = ? AND fiscal_sign = ?",
			receipt.FiscalNumber,
			receipt.FiscalDocument,
			receipt.FiscalSign,
		).Take(&existingReceipt).Error

		if err == nil {
			// Дубликат найден
			rest.WriteJSON(w, http.StatusConflict, rest.Response{
				Ok:     false,
				Error:  "Receipt already exists",
				Result: existingReceipt.Id,
			})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Ошибка при проверке
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		err = db.Create(&receipt).Error
		if err != nil {
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok:     true,
			Result: receipt.Id,
		})
	})
}

// GetReceiptByID возвращает чек по его Id
// @Summary Get receipt by Id
// @Description Retrieve a receipt by its unique Id.
// @Tags Receipt
// @Accept json
// @Produce json
// @Param id path string true "Receipt Id"
// @Success 200 {object} models.Receipt
// @Failure 404 {object} models.ErrorResponse "Чек не найден"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера"
// @Router /receipt/{id} [get]
func GetReceiptByID(db *gorm.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}

		var receipt models.Receipt
		err = db.Take(&receipt, id).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				rest.WriteError(w, http.StatusNotFound, errors.New("not found"))
				return
			}
			rest.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		rest.WriteJSON(w, http.StatusOK, rest.Response{
			Ok:     true,
			Result: receipt,
		})
	})
}
