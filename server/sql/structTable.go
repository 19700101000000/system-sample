package sql

type Employee struct {
	Id   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
