package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func DisplayOrdersTable(c echo.Context) error {

	strNum := echo.Context.Param(c, "num")

	num, err := strconv.ParseInt(strNum, 0, 64)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	listDatas := []ListData{}

	for i := int64(0); i < num; i++ {

		listDatas = append(
			listDatas,
			struct {
				ID       int64  `json:"id"`
				Client   User   `json:"client"`
				Employee User   `json:"employee"`
				Date     string `json:"date"`
			}{
				ID:       i,
				Client:   User{ID: "クライアントID", Name: "クライアント氏名"},
				Employee: User{ID: "社員ID", Name: "社員氏名"},
				Date:     "2018年11月19日",
			},
		)
	}

	return c.JSON(http.StatusOK, listDatas)
}
