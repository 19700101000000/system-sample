package handler

type AuthData struct {
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}
type AuthResponse struct {
	Success bool     `json:"success"`
	Auth    AuthData `json:"auth"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DisplayOrdersTable struct {
	BuyOrdersId  string `json:"ordersid"`
	ClientId     string `json:"clientid"`
	ClientName   string `json:"clientname"`
	EmployeeId   string `json:"employeeid"`
	EmployeeName string `json:"employeename"`
	InsertDate   string `json:"insertdate"`
}
