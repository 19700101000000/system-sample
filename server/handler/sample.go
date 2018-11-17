package handler

import (
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
