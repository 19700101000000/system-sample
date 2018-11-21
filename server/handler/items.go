package handler

import (
	"github.com/ShikinamiAsuka/ih13/server/sql"
	"github.com/labstack/echo"
	"net/http"
)

func ItemClients(c echo.Context) error {
	db, err := sql.NewDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	clients := sql.SelectClients(db)
	return c.JSON(http.StatusOK, clients)
}
func ItemEmployees(c echo.Context) error {
	db, err := sql.NewDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	employees := sql.SelectEmployees(db)
	return c.JSON(http.StatusOK, employees)
}
