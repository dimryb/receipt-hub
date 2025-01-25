package handlers_test

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"net/http"
	"receipt-loader/internal/models"
	"receipt-loader/internal/tests"
	"testing"
	"time"
)

func TestAddReceipt_Success(t *testing.T) {
	t.Log("TestAddReceipt_Success started")
	defer t.Log("TestAddReceipt_Success ended")
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	testTime := time.Now()
	receipt := models.Receipt{
		FiscalNumber:   123456789,
		FiscalDocument: 987654321,
		FiscalSign:     111222333,
		Date:           testTime.Format("2006-01-02"),
		Time:           testTime.Format("15:04:05"),
		Amount:         2,
	}

	receiptJSON, err := json.Marshal(receipt)
	assert.Nil(t, err)

	apitest.
		Handler(app.Router).
		Post("/receipt").
		JSON(receiptJSON).
		Expect(t).
		Status(http.StatusCreated).
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				NotEqual("result", nil). // Проверяем, что result содержит ID
				End(),
		).
		End()
}

func TestAddReceipt_Duplicate(t *testing.T) {
	t.Log("TestAddReceipt_Duplicate started")
	defer t.Log("TestAddReceipt_Duplicate ended")
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	// Добавляем первый чек
	testTime := time.Now()
	receipt := models.Receipt{
		FiscalNumber:   123456789,
		FiscalDocument: 987654321,
		FiscalSign:     111222333,
		Date:           testTime.Format("2006-01-02"),
		Time:           testTime.Format("15:04:05"),
		Amount:         2,
	}

	receiptJSON, err := json.Marshal(receipt)
	assert.Nil(t, err)

	// Создаем первый чек
	apitest.
		Handler(app.Router).
		Post("/receipt").
		JSON(receiptJSON).
		Expect(t).
		Status(http.StatusCreated).
		End()

	// Пытаемся добавить дубликат
	apitest.
		Handler(app.Router).
		Post("/receipt").
		JSON(receiptJSON).
		Expect(t).
		Status(http.StatusConflict).
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", false).
				Equal("error", "Receipt already exists").
				NotEqual("result", nil).
				End(),
		).
		End()
}

func TestGetReceiptByID(t *testing.T) {
	t.Log("TestGetReceiptByID started")
	defer t.Log("TestGetReceiptByID ended")

	app := tests.AppSetup(t) // Настраиваем приложение
	defer tests.AppTeardown(app)

	// Создаем тестовую запись в базе данных
	testReceipt := models.Receipt{
		Date:           "2025-01-25",
		Time:           "12:34:56",
		Amount:         123.45,
		FiscalNumber:   123456789,
		FiscalDocument: 987654321,
		FiscalSign:     111222333,
	}
	err := app.DB.Create(&testReceipt).Error
	if err != nil {
		t.Fatalf("failed to create test receipt: %v", err)
	}

	t.Run("Success", func(t *testing.T) {
		apitest.
			Handler(app.Router).
			Get(fmt.Sprintf("/receipt/%d", testReceipt.Id)).
			Expect(t).
			Status(http.StatusOK).
			Header("Content-Type", "application/json").
			Assert(
				jsonpath.Chain().
					Equal("ok", true).
					Equal("result.id", float64(testReceipt.Id)). // ID в JSON передается как float64
					Equal("result.fiscal_number", float64(testReceipt.FiscalNumber)).
					End(),
			).
			End()
	})

	t.Run("NotFound", func(t *testing.T) {
		apitest.
			Handler(app.Router).
			Get("/receipt/99999"). // Несуществующий ID
			Expect(t).
			Status(http.StatusNotFound).
			Header("Content-Type", "application/json").
			Assert(
				jsonpath.Chain().
					Equal("ok", false).
					Equal("error", "not found").
					End(),
			).
			End()
	})

	t.Run("InvalidID", func(t *testing.T) {
		apitest.
			Handler(app.Router).
			Get("/receipt/invalid-id"). // Некорректный ID
			Expect(t).
			Status(http.StatusNotFound).
			Header("Content-Type", "application/json").
			Assert(
				jsonpath.Chain().
					Equal("ok", false).
					Present("error").
					End(),
			).
			End()
	})
}
