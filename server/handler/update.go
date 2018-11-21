package handler

import (
	"github.com/ShikinamiAsuka/ih13/server/sql"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func UpdateOrder(c echo.Context) error {
	time.Sleep(1 * time.Second)
	db, err := sql.NewDB()
	if err != nil {
		return c.JSON(http.StatusOK, UpdateResult{Result: false})
	}
	defer db.Close()

	return c.JSON(http.StatusOK, UpdateResult{Result: true})
}
