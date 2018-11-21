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

//受注・見積りリスト
func SelectOrderList(db *gorm.DB) []Buy_order {
	orderlist := []Buy_order{}

	query := db.Table("buy_orders bo").
		Select("bo.id AS buy_orders_id,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,bo.insert_date AS insert_date,WEEKDAY(bo.insert_date) AS weekday").
		Joins("INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id").
		Where("bo.estimate_flag = ? ", 1)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		bOrder := Buy_order{}
		err := query.ScanRows(rows, &bOrder)
		if err != nil {
			log.Fatal(err)
		}
		orderlist = append(orderlist, bOrder)
	}
	return orderlist
}

//仕入リスト
func SelectPurchaseList(db *gorm.DB) []Buy_order {
	purchaseList := []Buy_order{}

	query := db.Table("buy_orders bo").
		Select("bo.id AS buy_orders_id,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,bo.insert_date AS insert_date,WEEKDAY(bo.insert_date) AS weekday").
		Joins("INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id").
		Where("bo.estimate_flag = ?", 1)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		bOrder := Buy_order{}
		err := query.ScanRows(rows, &bOrder)
		if err != nil {
			log.Fatal(err)
		}
		purchaseList = append(purchaseList, bOrder)
	}
	return purchaseList
}

func SelectClients(db *gorm.DB) []Client {
	clients := []Client{}

	db.SingularTable(true)
	db.Find(&clients)

	return clients
}
