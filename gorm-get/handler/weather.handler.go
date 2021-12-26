package handler

import (
	"net/http"

	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

type WeatherHandler interface {
	GetWeather(ctx *gin.Context)
}

type weatherHandler struct {
	weatherService service.WeatherService
}

func NewWeatherHandler(weatherService service.WeatherService) WeatherHandler {
	return &weatherHandler{
		weatherService: weatherService,
	}
}

func (c *weatherHandler) GetWeather(contextGin *gin.Context) {
	city := contextGin.Query("city")
	data, err := c.weatherService.GetWeatherFromDB(city)
	if err != nil {
		contextGin.IndentedJSON(http.StatusNotFound, gin.H{"message %s": err.Error()})
	} else {
		contextGin.JSON(http.StatusOK, data)
	}
}
