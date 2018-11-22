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
type BuyPurchaseList struct {
	Id           int64     `json:"id"`
	Manufacturer string    `json:"manufacturer"`
	Carname      string    `json:"carname"`
	Carmodelyear string    `json:"carmodelyear"`
	Budget       string    `json:"budget"`
	Employee     User      `json:"employee"`
	Client       User      `json:"client"`
	Date         time.Time `json:"date"`
}
type BuyClaimList struct {
	Id       int64     `json:"id"`
	Employee User      `json:"employee"`
	Client   User      `json:"client"`
	Date     time.Time `json:"date"`
}
type BuyRecoveryList struct {
	Id       int64     `json:"id"`
	Employee User      `json:"employee"`
	Client   User      `json:"client"`
	Date     time.Time `json:"date"`
}

type ListRequest struct {
	ListType int64 `json:"listtype"`
	ListFlag int64 `json:"listflag"`
}
type Color struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}

type UpdateResult struct {
	Result bool `json:"result"`
}
