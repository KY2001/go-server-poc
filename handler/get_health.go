package handler

import (
	"net/http"
	"strings"

	cloudsql "github.com/KY2001/go-server-poc/infrastructure/db/cloud-sql"
	"github.com/KY2001/go-server-poc/openapi"

	"github.com/labstack/echo/v4"
)

type GetHealthHandler struct {
}

func (h *GetHealthHandler) GetHealth(ctx echo.Context) error {
	err := checkDBHealth()
	if err != nil {
		return err
	}

	res := openapi.GetHealthResponse{
		Status: "OK",
	}
	return ctx.JSON(http.StatusOK, res)
}

func checkDBHealth() error {
	db := cloudsql.GetClient()

	insertHealth := replaceDoubleQuotesToBackTicks(
		`INSERT INTO ”health” (”message”) VALUES ('OK');`)
	_, err := db.Exec(insertHealth)
	if err != nil {
		return err
	}

	var (
		helthId   int
		message   string
		createdAt string
		updatedAt string
	)
	selectHealth := replaceDoubleQuotesToBackTicks(
		`SELECT * FROM ”health”;`)
	err = db.QueryRow(selectHealth).Scan(&helthId, &message, &createdAt, &updatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Backticks (“) are not allowed to be contained in a backtick string.
// So, we use ” instead and later, replace ” to backticks.
// See: https://stackoverflow.com/questions/21198980/how-to-escape-back-ticks#:~:text=I%20just%20used%20a%20placeholder
func replaceDoubleQuotesToBackTicks(s string) string {
	return strings.ReplaceAll(s, "”", "`")
}
