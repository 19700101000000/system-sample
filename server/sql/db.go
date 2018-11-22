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
func SelectOrderList(db *gorm.DB) []BuyOrderListItem {
	orderlist := []BuyOrderListItem{}

	query := db.Table("buy_orders bo").
		Select("bo.id AS buy_orders_id,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,bo.insert_date AS insert_date,WEEKDAY(bo.insert_date) AS weekday").
		Joins("INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id").
		Where("bo.estimate_flag = ? ", 1)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		bOrder := BuyOrderListItem{}
		err := query.ScanRows(rows, &bOrder)
		if err != nil {
			log.Fatal(err)
		}
		orderlist = append(orderlist, bOrder)
	}
	return orderlist
}

//仕入リスト
func SelectPurchaseList(db *gorm.DB) []BuyPurchaseListItem {
	purchaseList := []BuyPurchaseListItem{}

	query := db.Table("buy_purchase bp").
		Select("bp.id AS buy_purchase_id,bo.manufacturer AS manufacturer,bo.car_name AS carname,bo.car_model_year AS carmodelyear,bo.car_budget AS budget,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,bo.insert_date AS date,WEEKDAY(bo.insert_date) AS weekday,bp.delivery_date AS delivery_date").
		Joins("INNER JOIN buy_orders bo ON bp.buy_orders_id=bo.id INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id INNER JOIN suppliers s ON bp.suppliers_id=s.id").
		Where("bp.purchase_availability = ?", 1)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		bPurchase := BuyPurchaseListItem{}
		err := query.ScanRows(rows, &bPurchase)
		if err != nil {
			log.Fatal(err)
		}
		purchaseList = append(purchaseList, bPurchase)
	}
	return purchaseList
}

//請求リスト
func SelectClaimList(db *gorm.DB) []BuyClaimListItem {
	claimList := []BuyClaimListItem{}

	query := db.Table("buy_claim bc").
		Select("bc.id AS buy_claim_id,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,bo.insert_date AS date,bc.deadline AS deadline,WEEKDAY(bo.insert_date) AS weekday").
		Joins("INNER JOIN buy_purchase bp ON bc.buy_purchase_id=bp.id INNER JOIN buy_orders bo ON bp.buy_orders_id=bo.id INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id").
		Where("bc.billed = ?", 0)
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		bClaim := BuyClaimListItem{}
		err := query.ScanRows(rows, &bClaim)
		if err != nil {
			log.Fatal(err)
		}
		claimList = append(claimList, bClaim)
	}
	return claimList
}

//回収リスト
func SelectRecoveryList(db *gorm.DB) []BuyRecoveryListItem {
	recoveryList := []BuyRecoveryListItem{}

	query := db.Table("buy_recovery br").
		Select("br.id AS buy_recovery_id,bo.client_id,c.name AS client_name,bo.employee_id,e.name AS employee_name,bo.insert_date AS date,br.recovery_date AS recovery_date,WEEKDAY(bo.insert_date) AS weekday").
		Joins("INNER JOIN buy_claim bc ON br.buy_claim_id = bc.id INNER JOIN buy_purchase bp ON bc.buy_purchase_id=bp.id INNER JOIN buy_orders bo ON bp.buy_orders_id=bo.id INNER JOIN client c ON bo.client_id=c.id INNER JOIN employee e ON bo.employee_id=e.id").
		Where("")
	rows, err := query.Rows()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		bRecovery := BuyRecoveryListItem{}
		err := query.ScanRows(rows, &bRecovery)
		if err != nil {
			log.Fatal(err)
		}
		recoveryList = append(recoveryList, bRecovery)
	}
	return recoveryList
}

func SelectClients(db *gorm.DB) []Client {
	clients := []Client{}

	db.SingularTable(true)
	db.Find(&clients)

	return clients
}
