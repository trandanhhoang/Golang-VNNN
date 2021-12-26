package main

import (
	"fmt"
	"log"
	"os"

	"example/web-service-gin-channel/config"
	"example/web-service-gin-channel/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	weatherHandler handler.WeatherHandler = handler.NewWeatherHandler()
)

func main() {
	config.SetupDatabaseConnectionGORM()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	router := gin.Default()

	router.POST("/report_weather", weatherHandler.PostWeather)

	fmt.Println("Server is running on port " + port)
	router.Run("localhost:" + port)
}
