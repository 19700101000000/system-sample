package sql

type Employee struct {
	Id   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

type Buy_order struct {
	Buy_orders_id *string `json:"id"`
	Client_id     string  `json:"clientid"`
	Client_name   string  `json:"clientname"`
	Employee_id   string  `json:"employeeid"`
	Employee_name string  `json:"employeename"`
	Insert_date   string  `json:"date"`
}
