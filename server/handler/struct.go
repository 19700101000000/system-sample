package handler

import "time"

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

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type BuyOrderList struct {
	Id       int64     `json:"id"`
	Employee User      `json:"employee"`
	Client   User      `json:"client"`
	Date     time.Time `json:"date"`
}
type ListRequest struct {
	ListType int64 `json:"listtype"`
	ListFlag int64 `json:"listflag"`
}
