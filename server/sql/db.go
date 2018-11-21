package sql

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "server:server@tcp(db:3306)/ih2018_db?charset=utf8&parseTime=True&loc=Local")
}

func SelectEmployees(db *gorm.DB) []Employee {
	employees := []Employee{}

	/*テーブルが複数形出ない場合に宣言*/
	db.SingularTable(true)
	db.Find(&employees)

	return employees
}

func SelectOrderList(db *gorm.DB) []Buy_order {
	orderlist := []Buy_order{}

	query := db.Table("buy_orders bo").
		Select("bo.id AS buy_orders_id,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,DATE_FORMAT(bo.insert_date,'%Y年%m月%d日 %k時%i分') AS insert_date,WEEKDAY(bo.insert_date) AS weekday").
		Joins("INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id").
		Where("bo.estimate_flag = ?", 1)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}

	var Buy_order struct {
		Buy_order
		Buy_orders_id *string `json:"id"`
		Client_id     string  `json:"clientid"`
		Client_name   string  `json:"clientname"`
		Employee_id   string  `json:"employeeid"`
		Employee_name string  `json:"employeename"`
		Insert_date   string  `json:"date"`
	}

	var orderList []string
	for rows.Next() {
		err := query.ScanRows(rows, &Buy_order)
		if err != nil {
			log.Fatal(err)
		}
		orderList = append(orderList, *Buy_order.Buy_orders_id)
		orderList = append(orderList, Buy_order.Client_id)
		orderList = append(orderList, Buy_order.Client_name)
		orderList = append(orderList, Buy_order.Employee_id)
		orderList = append(orderList, Buy_order.Employee_name)
		orderList = append(orderList, Buy_order.Insert_date)
	}

	/*テーブルが複数形出ない場合に宣言*/
	// db.SingularTable(true)
	//db.Find(&orderlist)

	return orderlist
}
