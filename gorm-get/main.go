package main

import (
	"fmt"
	"log"
	"os"

	"example/web-service-gin/config"
	"example/web-service-gin/handler"
	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	weatherService service.WeatherService = service.NewWeatherService()
	weatherHandler handler.WeatherHandler = handler.NewWeatherHandler(weatherService)
)

func main() {
	config.SetupDatabaseConnectionGORM()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/weather", weatherHandler.GetWeather)

	fmt.Println("Server is running on port " + port)
	router.Run("localhost:" + port)
}
