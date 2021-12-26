package thread

import (
	"errors"
	"example/web-service-gin-channel/config"
	_weather "example/web-service-gin-channel/models"
	"sync"

	"gorm.io/gorm"
)

func HandleMessages(channel chan _weather.Weather, wg *sync.WaitGroup, fatalErrors chan error, successAddToDB chan bool) {
	select {
	case msg := <-channel:
		if msg.DataInNext12H == "0" {
			fatalErrors <- errors.New("Fake Error!!!")
		} else {
			if err := storeInDatabase(msg); err != nil {
				fatalErrors <- err
			} else {
				fatalErrors <- nil
			}
		}
	}
	wg.Done()
}

func storeInDatabase(data _weather.Weather) error {
	var db *gorm.DB = config.SetupDatabaseConnectionGORM()
	defer config.CloseConnectionGORM(db)

	if err := db.Select("name", "data").Create(&data).Error; err != nil {
		return errors.New("Can't add record weather into database")
	}
	return nil
}
