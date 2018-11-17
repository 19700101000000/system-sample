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

type DisplayTable struct {
	Orderid string `json:"id"`
}
