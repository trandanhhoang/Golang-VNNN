package db

import (
	"fmt"
	"os"

	log "user-service/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnectionGORM() *gorm.DB {
	if err := godotenv.Load("../.env"); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error loading .env file")
	}

	port := os.Getenv("USER_MYSQL_URL")

	dsn := fmt.Sprintf("admin:password@tcp(%s)/user?charset=utf8mb4&parseTime=True&loc=Local", port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func CloseConnectionGORM(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
