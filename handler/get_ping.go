package handler

import (
	"net/http"

	"github.com/KY2001/go-server-poc/openapi"

	"github.com/labstack/echo/v4"
)

type GetPingHandler struct {
}

func (h *GetPingHandler) GetPing(ctx echo.Context) error {
	res := openapi.GetPingResponse{
		Status: "OK",
	}
	return ctx.JSON(http.StatusOK, res)
}
