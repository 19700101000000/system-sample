package handler

import (
	"github.com/ShikinamiAsuka/ih13/server/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Sample(c echo.Context) error {
	strNum := echo.Context.Param(c, "num")

	num, err := strconv.ParseInt(strNum, 0, 64)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var responseJson []struct {
		Id   int64
		Text string
	}

	for i := int64(0); i < num; i++ {
		responseJson = append(
			responseJson,
			struct {
				Id   int64
				Text string
			}{
				Id:   i,
				Text: string('a' + i%26),
			},
		)
	}

	return c.JSON(http.StatusOK, responseJson)
}

func SampleDB(c echo.Context) error {
	db, err := sql.NewDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	employees := sql.SelectEmployees(db)
	return c.JSON(http.StatusOK, employees)
}
