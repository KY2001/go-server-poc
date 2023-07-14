package handler

import (
	"net/http"

	"github.com/KY2001/go-server-poc/openapi"

	"github.com/labstack/echo/v4"
)

type GetHealthHandler struct {
}

func (h *GetHealthHandler) GetHealth(ctx echo.Context) error {
	// TODO: check the DB connection
	// if there is a problem, return error

	res := openapi.GetHealthResponse{
		Status: "OK",
	}

	return ctx.JSON(http.StatusOK, res)
}
