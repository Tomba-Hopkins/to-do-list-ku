package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/todo?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Database connected")


	DB = db
}