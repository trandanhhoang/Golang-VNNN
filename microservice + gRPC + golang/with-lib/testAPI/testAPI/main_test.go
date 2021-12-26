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

var IDLENGTH = 10

var testCases = []struct {
	quantity int
}{
	// {10},
	{1},
}

func TestPostOrder(t *testing.T) {
	c := client.NewClient(
		"http://localhost:8888/order",
	)
	for _, test := range testCases {

		orderFile, err := os.OpenFile("../mock-data/order.json", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer orderFile.Close()

		orderRequest := models.OrderRequest{}

		orderFromFile, err := utils.ReadFileJSON(orderFile)
		err = json.Unmarshal(orderFromFile, &orderRequest)
		orderRequest.Products[0].Quantity = test.quantity
		res, err := c.CallGhtkApi(&orderRequest, http.MethodPost)
		if err != nil {
			fmt.Println(err)
		}

		result, err := json.Marshal(res)
		fmt.Println(string(result))
		fmt.Println("---------------------------")
	}
}
