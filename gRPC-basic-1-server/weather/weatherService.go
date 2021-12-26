package weather

import (
	"context"
	"log"
)

type WeatherService struct{}

func (s *WeatherService) GetWeather(ctx context.Context, req *WeatherRequest) (*WeatherRespond, error) {
	log.Println("Received mess", req.Name)
	// Now it's a fixed Data. I will take data from database soon
	return &WeatherRespond{
		Name: "hoang",
		Data: "cloudy",
	}, nil
}
