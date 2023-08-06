package handler

import (
	"fmt"
	"net/http"

	cloudsql "github.com/KY2001/go-server-poc/infrastructure/db/cloud-sql"
	"github.com/KY2001/go-server-poc/openapi"

	"github.com/labstack/echo/v4"
)

type GetUserHandler struct {
}

func (h *GetUserHandler) GetUser(ctx echo.Context, userId string) error {
	userName, err := getUserName(userId)
	if err != nil {
		return err
	}

	res := openapi.GetUserResponse{
		UserName: userName,
	}
	return ctx.JSON(http.StatusOK, res)
}

func getUserName(userId string) (string, error) {
	db := cloudsql.GetClient()

	var userName string
	selectUserName := replaceDoubleQuotesToBackTicks(
		fmt.Sprintf(`SELECT ”user_name” FROM ”user” WHERE ”user_id” = %q;`, userId))
	err := db.QueryRow(selectUserName).Scan(&userName)
	if err != nil {
		return "", err
	}

	return userName, nil
}
