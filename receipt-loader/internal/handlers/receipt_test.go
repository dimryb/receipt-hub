package handlers_test

import (
	"encoding/json"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"receipt-loader/internal/models"
	"receipt-loader/internal/tests"
	"testing"
	"time"
)

func TestDatabaseConnection(t *testing.T) {
	dsn := "postgres://postgres@localhost:5432/receipts_test"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Exec("SELECT 1").Error
	assert.Nil(t, err, "Database is not ready: %v", err)

	t.Log("Database connection is ready and working.")
}

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
