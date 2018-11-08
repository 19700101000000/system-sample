package handler

type AuthData struct {
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}
type AuthResponse struct {
	Success bool     `json:"success"`
	Auth    AuthData `json:"auth"`
}
