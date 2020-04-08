package configs

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// MysqlClient is a global variable.
// Use this variable instead of repeating calling ConnectMysql
var MysqlClient *sql.DB

// ConnectMysql called in main.go
func ConnectMysql() {
	sqlConfig := &mysql.Config{
		Net:    "tcp",
		User:   "kwantz",
		Passwd: "kwantz123",
		DBName: "db_crud",
		Addr:   "crud-mysql:3306",
	}

	log.Print("Connecting MySQL ... ")
	client, err := sql.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		log.Println("Error")
		log.Fatal(err.Error())
	}

	log.Println("Success")
	MysqlClient = client
}
