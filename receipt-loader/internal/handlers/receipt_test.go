package handlers_test

import (
	"encoding/json"
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
