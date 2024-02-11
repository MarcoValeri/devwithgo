package util

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	psh "github.com/platformsh/config-reader-go/v2"
	sqldsn "github.com/platformsh/config-reader-go/v2/sqldsn"
)

var dbPlatformSh *sql.DB
var dbLocal *sql.DB

func DatabaseConnectionPlatformSh() *sql.DB {
	config, err := psh.NewRuntimeConfig()
	if err != nil {
		log.Fatal("Some error occured. Err: $s", err)
	}

	credentials, err := config.Credentials("database")
	if err != nil {
		log.Fatal("Some error occured. Err: $s", err)
	}

	formatted, err := sqldsn.FormattedCredentials(credentials)
	if err != nil {
		log.Fatal("Some error occured. Err: $s", err)
	}

	dbPlatformSh, err := sql.Open("mysql", formatted)
	if err != nil {
		log.Fatal("Some error occured. Err: $s", err)
	}

	return dbPlatformSh
}

func DatabaseLocalConnection() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Some error occured. Err: $s", err)
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
