package db

import (
	"crypto/sha256"
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

func Auth(name, pass string) *string {
	var result string

	err := db.QueryRow(
		"SELECT `name` FROM `user` WHERE `name` = ? AND `password` = ? AND alive = 1",
		name,
		fmt.Sprintf("%x", sha256.Sum256([]byte(pass))),
	).Scan(&result)
	if err != nil {
		fmt.Printf("error by db.Auth:: %v\n", err)
		return nil
	}
	return &result
}
