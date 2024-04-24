package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	psh "github.com/platformsh/config-reader-go/v2"
	sqldsn "github.com/platformsh/config-reader-go/v2/sqldsn"
)

var dbPlatformSh *sql.DB
var dbLocal *sql.DB

func DatabaseConnection() *sql.DB {

	flag := "local"

	if flag == "local" {
		return databaseConnectionLocal()
	}

	if flag == "platform" {
		return databaseConnectionPlatformSh()
	}

	return databaseConnectionPlatformSh()
}

func databaseConnectionPlatformSh() *sql.DB {
	config, err := psh.NewRuntimeConfig()
	if err != nil {
		fmt.Println("Some error occured. Err: $s", err)
	}

	credentials, err := config.Credentials("database")
	if err != nil {
		fmt.Println("Some error occured. Err: $s", err)
	}

	formatted, err := sqldsn.FormattedCredentials(credentials)
	if err != nil {
		fmt.Println("Some error occured. Err: $s", err)
	}

	dbPlatformSh, err := sql.Open("mysql", formatted)
	if err != nil {
		fmt.Println("Some error occured. Err: $s", err)
	}

	return dbPlatformSh
}

func databaseConnectionLocal() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Some error occured. Err: $s", err)
	}

	dbCredentials := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_USER_PW"),
		Net:                  os.Getenv("DB_USER_NET"),
		Addr:                 os.Getenv("DB_USER_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	// Get a database handle
	dbLocal, err = sql.Open("mysql", dbCredentials.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pinErr := dbLocal.Ping()
	if pinErr != nil {
		log.Fatal(pinErr)
	}

	return dbLocal
}
