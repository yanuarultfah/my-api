package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Client      *sql.DB
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "postgres"
	DB_SSLMODE  = "disable"
)

func init() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE)

	var err error
	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
