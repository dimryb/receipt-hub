package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"receipt-loader/internal/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	handlers.Ping(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "OK", w.Body.String())
}
