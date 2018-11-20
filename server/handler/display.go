package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func DisplayOrdersTable(c echo.Context) error {

	strNum := echo.Context.Param(c, "listtype")
	num, err := strconv.ParseInt(strNum, 0, 64)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	listType := new(ListRequest)
	if err := c.Bind(listType); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var test = ""
	switch num {
	case 0:
		test = "見積もり"
		break
	case 1:
		test = "受注"
		break
	case 2:
		test = "仕入"
		break
	case 3:
		test = "出荷"
		break
	case 4:
		test = "請求"
		break
	case 5:
		test = "回収"
		break
	}

	listDatas := []ListData{}

	for i := int64(0); i < 5; i++ {

		listDatas = append(
			listDatas,
			struct {
				ID       int64  `json:"id"`
				Client   User   `json:"client"`
				Employee User   `json:"employee"`
				Date     string `json:"date"`
			}{
				ID:       1,
				Client:   User{ID: "999999", Name: test},
				Employee: User{ID: "99999", Name: "社員氏名"},
				Date:     "2018年11月19日",
			},
		)
	}

	return c.JSON(http.StatusOK, listDatas)
}
