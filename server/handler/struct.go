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

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ListData struct {
	ID       int64  `json:"id"`
	Client   User   `json:"client"`
	Employee User   `json:"employee"`
	Date     string `json:"date"`
}
