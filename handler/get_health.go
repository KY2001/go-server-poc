package handler

import (
	"net/http"

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
	db, err := cloudsql.GetDB()
	if err != nil {
		return err
	}

	createHealth := `CREATE TABLE IF NOT EXISTS health (
		id SERIAL NOT NULL,
		created_at datetime NOT NULL,
		message VARCHAR(6) NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = db.Exec(createHealth)
	if err != nil {
		return err
	}

	insertHealth := `INSERT INTO health (created_at, message) VALUES (NOW(), 'OK');`
	_, err = db.Exec(insertHealth)
	if err != nil {
		return err
	}

	var (
		id        int
		createdAt string
		message   string
	)
	selectHealth := `SELECT * FROM health;`
	err = db.QueryRow(selectHealth).Scan(&id, &createdAt, &message)
	if err != nil {
		return err
	}

	return nil
}
