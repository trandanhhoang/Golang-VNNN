package ghtk

import (
	"encoding/json"
	"fmt"
	client "ghtk/client"
	"ghtk/models"
	"ghtk/utils"
	"net/http"
	"os"
	"testing"
)

var testCases = []string{
	"2-chart.json",
	// "complicated-chart.json",
	// "pie-chart.json",
}

func TestPostOrder(t *testing.T) {
	c := client.NewClient(
		"http://localhost:3000/api/chart/demo",
	)
	for _, test := range testCases {

		orderFile, err := os.OpenFile("../mock-data/"+test, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer orderFile.Close()

		orderRequest := models.OrderRequest{}

		orderFromFile, err := utils.ReadFileJSON(orderFile)
		err = json.Unmarshal(orderFromFile, &orderRequest)

		res, err := c.CallGhtkApi(&orderRequest, http.MethodPost)
		if err != nil {
			fmt.Println(err)
		}

		result, err := json.Marshal(res)
		fmt.Println(string(result))
		fmt.Println("---------------------------")
	}
}
