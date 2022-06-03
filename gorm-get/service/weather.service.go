package service

import (
	"database/sql"
	"errors"
	"example/web-service-gin/config"
	_weather "example/web-service-gin/service/weather"

	"gorm.io/gorm"
)

type WeatherService interface {
	GetWeatherFromDB(city string) (*_weather.Weather, error)
}

type weatherService struct {
	db *sql.DB
}

func NewWeatherService() WeatherService {
	return &weatherService{}
}

func (c *weatherService) GetWeatherFromDB(city string) (*_weather.Weather, error) {

	var weather _weather.Weather
	var tbl _weather.Tabler = weather

	var db *gorm.DB = config.SetupDatabaseConnectionGORM()

	defer config.CloseConnectionGORM(db)

	if err := db.Table(tbl.TableName()).First(&weather, "name = ?", city).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Can't find name of city in database")
		} else {
			return nil, err
		}
	}
	return &weather, nil
}
