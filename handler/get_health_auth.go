package handler

import (
	"net/http"

	"github.com/KY2001/go-server-poc/openapi"

	"github.com/labstack/echo/v4"
)

type GetHealthAuthHandler struct {
}

func (h *GetHealthAuthHandler) GetHealthAuth(ctx echo.Context) error {
	res := openapi.GetHealthAuthResponse{
		Status: "OK",
	}
	return ctx.JSON(http.StatusOK, res)
}
