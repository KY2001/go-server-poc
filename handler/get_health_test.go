package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KY2001/go-server-poc/handler"
	cloudsql "github.com/KY2001/go-server-poc/infrastructure/db/cloud-sql"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetHealthHandler_GetHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)

	cloudsql.InitClient()

	GetHealthHandler := handler.GetHealthHandler{}
	err := GetHealthHandler.GetHealth(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"status": "OK"}`, rec.Body.String())
}
