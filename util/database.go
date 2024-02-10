package util

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseConnection() *sql.DB {
	dbCredentials := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "devwithgo",
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", dbCredentials.FormatDSN())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB credential: TRUE")
	}

	pinErr := db.Ping()
	if pinErr != nil {
		log.Fatal(pinErr)
	} else {
		fmt.Println("DB Ping: TRUE")
	}

	return db
}
