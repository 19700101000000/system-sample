package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func InitDB() {
	DB, err := SqlConnection()
	if err != nil {
		panic("DB connecion failed")
	}
	db = DB
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func SqlConnection() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	/* connection db */
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName))
}
