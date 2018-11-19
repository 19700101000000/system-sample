package handler

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "root:root@DB_HOST/ih2018_db")
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail db connection")
	}
	defer db.Close()
	return c.String(http.StatusOK, "success db connection")
}
