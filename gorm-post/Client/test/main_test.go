package weather

import (
	"encoding/json"
	"fmt"
	"testing"
	"weather/client"
	"weather/models"
)

var IDLENGTH = 10

func TestPostOrder(t *testing.T) {
	c := client.NewClient(
		"http://localhost:4444/report_weather",
	)

	orderRequest := models.Weather{
		Name:          "Kualalumpua",
		DataInNext12H: "Storming",
	}

	// WRAP API TEST
	res, err := c.PostOrder(&orderRequest)
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
