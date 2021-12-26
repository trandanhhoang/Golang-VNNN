package service

import (
	"database/sql"
	"errors"
	_weather "example/web-service-gin/service/weather"
	"fmt"
)

type WeatherService interface {
	GetWeatherFromDB(city string) (*_weather.Weather, error)
}

type weatherService struct {
	db *sql.DB
}

func NewWeatherService(db *sql.DB) WeatherService {
	return &weatherService{
		db: db,
	}
}

func (c *weatherService) GetWeatherFromDB(city string) (*_weather.Weather, error) {

	var weather _weather.Weather
	query := fmt.Sprintf("SELECT * FROM weathers where name ='%s';", city)

	if err := c.db.QueryRow(query).Scan(&weather.ID, &weather.Name, &weather.DataInNext12H); err != nil {
		if err := c.db.QueryRow(query).Scan(&weather.ID, &weather.Name, &weather.DataInNext12H); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("Can't find name of city in database")
			} else {
				return nil, err
			}
		}
	}

	return &weather, nil
}
