package sql

type Employee struct {
	Id   string `gorm:"primary_key" json:"value"`
	Name string `json:"text"`
}

type Client struct {
	Id   int    `gorm:"primary_key" json:"value"`
	Name string `json:"text"`
}
