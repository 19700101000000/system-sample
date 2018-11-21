package handler

import (
	"net/http"
	"strconv"

	"github.com/ShikinamiAsuka/ih13/server/sql"
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

	db, err := sql.NewDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	if num == 0 || num == 1 {
		//見積もり
		listDbDatas := sql.SelectOrderList(db)
		result := []BuyOrderList{}
		for _, data := range listDbDatas {
			d := BuyOrderList{}
			d.Id = data.Buy_orders_id
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Insert_date
			result = append(result, d)
		}
		return c.JSON(http.StatusOK, result)
	} else if num == 2 {
		//仕入
		listDbDatas := sql.SelectPurchaseList(db)
		result := []BuyOrderList{}
		for _, data := range listDbDatas {
			d := BuyOrderList{}
			d.Id = data.Buy_orders_id
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Insert_date
			result = append(result, d)
		}
		return c.JSON(http.StatusOK, result)
	} else if num == 3 {
		//出荷
		listDbDatas := sql.SelectOrderList(db)
		result := []BuyOrderList{}
		for _, data := range listDbDatas {
			d := BuyOrderList{}
			d.Id = data.Buy_orders_id
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Insert_date
			result = append(result, d)
		}
		return c.JSON(http.StatusOK, result)
	} else if num == 4 {
		//請求
		listDbDatas := sql.SelectOrderList(db)
		result := []BuyOrderList{}
		for _, data := range listDbDatas {
			d := BuyOrderList{}
			d.Id = data.Buy_orders_id
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Insert_date
			result = append(result, d)
		}
		return c.JSON(http.StatusOK, result)
	} else if num == 5 {
		//回収
		listDbDatas := sql.SelectOrderList(db)
		result := []BuyOrderList{}
		for _, data := range listDbDatas {
			d := BuyOrderList{}
			d.Id = data.Buy_orders_id
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Insert_date
			result = append(result, d)
		}
		return c.JSON(http.StatusOK, result)
	} else {
		listDbDatas := sql.SelectOrderList(db)
		result := []BuyOrderList{}
		for _, data := range listDbDatas {
			d := BuyOrderList{}
			d.Id = data.Buy_orders_id
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Insert_date
			result = append(result, d)
		}
		return c.JSON(http.StatusOK, result)
	}
}
