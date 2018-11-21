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
