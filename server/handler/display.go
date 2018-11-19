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
	// var listDatas []struct {
	// 	Id      int64 `json:"id"`
	// 	Clients []struct {
	// 		Id string `json:"id"`
	// 		Name string `json:"name"`
	// 	}
	// 	Employee []struct {oneData []struct{}} `json:"employee"`
	// 	Date string `json:"date"`
	// }

	for i := int64(0); i < num; i++ {

		// User1 := User{{ID: "", Name: ""}}

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
				Employee: User{ID: "", Name: ""},
				Date:     "2018年11月19日",
			},
		)
	}

	return c.JSON(http.StatusOK, listDatas)
}
