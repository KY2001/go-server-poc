package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	cloudsql "github.com/KY2001/go-server-poc/infrastructure/db/cloud-sql"
	"github.com/KY2001/go-server-poc/openapi"

	"github.com/labstack/echo/v4"
)

type PostAuthSignupHandler struct {
}

func (h *PostAuthSignupHandler) PostAuthSignup(ctx echo.Context, params openapi.PostAuthSignupParams) error {
	err := createUser(params.UserName, params.UserId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

func createUser(userName string, userId string) error {
	db := cloudsql.GetClient()

	var count int
	selectUserName := replaceDoubleQuotesToBackTicks(
		fmt.Sprintf(`SELECT COUNT(*) FROM ”user” WHERE ”user_id” = %q OR ”user_name” = %q;`, userId, userName))
	err := db.QueryRow(selectUserName).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		return fmt.Errorf("createUser: user already exists")
	}

	insertUser := replaceDoubleQuotesToBackTicks(
		fmt.Sprintf(`INSERT INTO ”user” (”user_id”, ”user_name”) VALUES (%q, %q);`, userId, userName))
	_, err = db.Exec(insertUser)
	if err != nil {
		return err
	}

	return nil
}
