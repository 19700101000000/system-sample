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
		//見積もり&受注
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
	} else if num == 2 || num == 3 {
		//仕入&出荷
		listDbDatas := sql.SelectPurchaseList(db)
		result := []BuyPurchaseList{}
		for _, data := range listDbDatas {
			d := BuyPurchaseList{}
			d.Id = data.Buy_purchase_id
			d.Manufacturer = data.Manufacturer
			d.Carname = data.Carname
			d.Carmodelyear = data.Carmodelyear
			d.Budget = data.Budget
			d.Client.Id = data.Client_id
			d.Client.Name = data.Client_name
			d.Employee.Id = data.Employee_id
			d.Employee.Name = data.Employee_name
			d.Date = data.Date
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

func SampleDBList(c echo.Context) error {

	db, err := sql.NewDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	//仕入
	listDbDatas := sql.SelectPurchaseList(db)
	result := []BuyPurchaseList{}
	for _, data := range listDbDatas {
		d := BuyPurchaseList{}
		d.Id = 1
		d.Manufacturer = data.Manufacturer
		d.Carname = data.Carname
		d.Carmodelyear = data.Carmodelyear
		d.Budget = data.Budget
		d.Client.Id = data.Client_id
		d.Client.Name = data.Client_name
		d.Employee.Id = data.Employee_id
		d.Employee.Name = data.Employee_name
		d.Date = data.Date
		result = append(result, d)
	}
	return c.JSON(http.StatusOK, result)

}
