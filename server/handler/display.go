package handler

import (
	"net/http"

	"github.com/labstack/echo"
)
func CreateOrderTable(c echo.Context) error {
	
	data := DisplayTable{
		Orderid: "orderid",
	}

	return c.JSON(http.StatusOK, data)
}
