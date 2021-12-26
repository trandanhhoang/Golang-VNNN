package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/weathers")
	if err != nil {
		panic(err.Error())
	}

	return db
}
