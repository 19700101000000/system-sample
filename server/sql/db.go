package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "server:server@tcp(db:3306)/ih2018_db?charset=utf8&parseTime=True&loc=Local")
}

func SelectEmployees(db *gorm.DB) []Employee {
	employees := []Employee{}

	/*テーブルが複数形出ない場合に宣言*/
	db.SingularTable(true)
	db.Find(&employees)

	return employees
}
