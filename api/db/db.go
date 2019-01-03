package db

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	arks "github.com/19700101000000/system-sample/api/ark"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DB struct {
	Sql    *sql.DB
	Stream chan interface{}
}

func NewDB() *DB {
	db, err := SqlConnection()
	if err != nil {
		panic("DB connecion failed")
	}
	return &DB{
		Sql:    db,
		Stream: make(chan interface{}),
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
func (db *DB) Monitor() {
	for {
		select {
		case ark := <-db.Stream:
			db.separateArk(ark)
		}
	}
	close(db.Stream)
}

func (db *DB) separateArk(ark interface{}) {
	switch v := ark.(type) {
	case arks.Auth:
		db.typeArkAuth(v)
	}
}
func (db *DB) typeArkAuth(ark arks.Auth) {
	defer close(ark.Result)
	var name *string
	err := db.Sql.QueryRow(
		"SELECT `name` FROM `user` WHERE `name` = ? AND `password` = ? AND alive = 1",
		ark.Name,
		fmt.Sprintf("%x", sha256.Sum256([]byte(ark.Pass))),
	).Scan(&name)
	if err != nil {
		fmt.Printf("error by typeArkAuth:: %v\n", err)
	}
	ark.Result <- name
}
