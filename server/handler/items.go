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
func ItemColors(c echo.Context) error {
	colors := []Color{
		{Value: 0, Text: "赤"},
		{Value: 1, Text: "橙"},
		{Value: 2, Text: "黄"},
		{Value: 3, Text: "黄緑"},
		{Value: 4, Text: "緑"},
		{Value: 5, Text: "水色"},
		{Value: 6, Text: "青"},
		{Value: 7, Text: "紫"},
		{Value: 8, Text: "白"},
		{Value: 9, Text: "黒"},
		{Value: 10, Text: "シルバー"},
		{Value: 11, Text: "ゴールド"},
		{Value: 12, Text: "その他"},
	}
	return c.JSON(http.StatusOK, colors)
}
