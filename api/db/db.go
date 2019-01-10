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

func Auth(user, pass string) (id int, name string, ok bool) {
	err := db.QueryRow(
		"SELECT `id`, `name` FROM `user` WHERE `name` = ? AND `password` = ? AND alive = 1",
		user,
		fmt.Sprintf("%x", sha256.Sum256([]byte(pass))),
	).Scan(&id, &name)
	if err != nil {
		fmt.Printf("error by db.Auth:: %v\n", err)
		ok = false
		return
	}
	ok = true
	return
}

func Categories() map[string]interface{} {
	result := make(map[string]interface{})

	rows, err := db.Query("SELECT `id`, `name` FROM `category`")
	if err != nil {
		fmt.Printf("error by db.GetCategories:: %v\n", err)
		return result
	}

	options := make([]OptionCategory, 0)
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Printf("error by db.GetCategories:: %v\n", err)
			continue
		}
		options = append(options, OptionCategory{Text: name, Value: id})
	}
	result["options"] = options
	return result
}
