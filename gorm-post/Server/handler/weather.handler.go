package handler

import (

	// "weather/models"

	_weather "example/web-service-gin-channel/models"
	"example/web-service-gin-channel/service/thread"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type WeatherHandler interface {
	PostWeather(ctx *gin.Context)
}

type weatherHandler struct {
}

func NewWeatherHandler() WeatherHandler {
	return &weatherHandler{}
}

// type Respond struct {
// 	RespondError error
// 	Success      bool
// }

func (c *weatherHandler) PostWeather(contextGin *gin.Context) {
	data := _weather.Weather{}
	contextGin.BindJSON(&data)

	channel := make(chan _weather.Weather)
	DbErrors := make(chan error)
	successAddToDB := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go thread.HandleMessages(channel, &wg, DbErrors, successAddToDB)

	channel <- data
	select {
	case err := <-DbErrors:
		close(DbErrors)
		if err != nil {
			respondWithError := _weather.WeatherBaseResponse{
				Message: err.Error(),
				Success: false,
			}
			contextGin.JSON(http.StatusNotFound, respondWithError)
		} else {
			respond := _weather.WeatherBaseResponse{
				Message: "Insert into table success",
				Success: true,
			}
			contextGin.JSON(http.StatusOK, respond)
		}
	}
	//Did I need to run it in another goroutine?
	fmt.Println("heck1")
	wg.Wait()
	close(channel)
	fmt.Println("heck2")

}
