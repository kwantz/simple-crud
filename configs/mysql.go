package configs

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// ConnectMySQL - MySQL connection
func ConnectMySQL() *sqlx.DB {
	dbUser := "kwantz"
	dbPass := "kwantz123"
	dbName := "db_crud"
	dbAddr := "crud-mysql:3306"

	option := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbAddr, dbName)
	conn, err := sqlx.Connect("mysql", option)
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}
