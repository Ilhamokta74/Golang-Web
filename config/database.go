package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-web?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db
}
