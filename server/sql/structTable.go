package sql

import "time"

type Employee struct {
	Id   string `gorm:"primary_key" json:"value"`
	Name string `json:"text"`
}

type Client struct {
	Id   int    `gorm:"primary_key" json:"value"`
	Name string `json:"text"`
}

type ClientUser struct {
	Client_id   string `json:"id"`
	Client_name string `json:"name"`
}
type EmployeeUser struct {
	Employee_id   string `json:"id"`
	Employee_name string `json:"name"`
}

type BuyOrderListItem struct {
	Buy_orders_id int64     `json:"id"`
	Client_id     string    `json:"client_id"`
	Client_name   string    `json:"client_name"`
	Employee_id   string    `json:"employee_id"`
	Employee_name string    `json:"employee_name"`
	Insert_date   time.Time `json:"date"`
}
type BuyPurchaseListItem struct {
	Buy_purchase_id int64     `json:"id"`
	Manufacturer    string    `json:"manufacturer"`
	Carname         string    `json:"carname"`
	Carmodelyear    string    `json:"carmodelyear"`
	Budget          string    `json:"budget"`
	Client_id       string    `json:"client_id"`
	Client_name     string    `json:"client_name"`
	Employee_id     string    `json:"employee_id"`
	Employee_name   string    `json:"employee_name"`
	Date            time.Time `json:"date"`
	Shipmentdate    time.Time `json:"shipmentdate"`
}
type BuyClaimListItem struct {
	Buy_claim_id  int64     `json:"id"`
	Client_id     string    `json:"client_id"`
	Client_name   string    `json:"client_name"`
	Employee_id   string    `json:"employee_id"`
	Employee_name string    `json:"employee_name"`
	Deadline      time.Time `json:"deadline"`
	Date          time.Time `json:"date"`
}
type BuyRecoveryListItem struct {
	Buy_recovery_id int64     `json:"id"`
	Client_id       string    `json:"client_id"`
	Client_name     string    `json:"client_name"`
	Employee_id     string    `json:"employee_id"`
	Employee_name   string    `json:"employee_name"`
	Recovery_date   time.Time `json:"recovery_date"`
	Insert_date     time.Time `json:"date"`
}
