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

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var IDLENGTH = 10

func TestPostOrder(t *testing.T) {
	c := client.NewClient(
		"https://services.ghtklab.com/services/shipment/order",
		"B5247cc50b15ab7C487494aD5e3F30aE41aa8bDd",
	)
	// CREATE DATA
	orderFile, err := os.OpenFile("../mock-data/order.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer orderFile.Close()

	idForOrder := utils.RandomId()
	orderFromFile, err := utils.ReadFileJSON(orderFile)
	orderRequest := models.GhtkOrderRequest{}
	err = json.Unmarshal(orderFromFile, &orderRequest)
	orderRequest.Order.ID = idForOrder

	//new Func
	// res, err := c.CallGhtkApi(&orderRequest)

	// WRAP API TEST

	// res, err := c.PostOrder(&orderRequest)
	res, err := c.CallGhtkApi(&orderRequest, http.MethodPost)
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.Marshal(res)
	fmt.Println(string(result))

	// WRITE INTO FILE "save-order.txt"
	f, err := os.OpenFile("save-order.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	orderJson, err := json.Marshal(res.Order)
	if err != nil {
		panic(err)
	}

	saveOrder := models.GhtkSaveOrder{}
	json.Unmarshal(orderJson, &saveOrder)

	saveOrderJson, err := json.Marshal(saveOrder)
	if err != nil {
		panic(err)
	}
	if _, err = f.WriteString(string(saveOrderJson) + "\n"); err != nil {
		panic(err)
	}
	assert.NotNil(t, res, "Respond not nil")
}

func TestGetShippingFee(t *testing.T) {
	// TEST Calculate shipping fee
	shippingFeeParams := map[string]interface{}{
		"pick_province":  "Hà Nội",
		"pick_district":  "Quận Hai Bà Trưng",
		"province":       "Hà nội",
		"district":       "Quận Cầu Giấy",
		"address":        "P.503 tòa nhà Auu Việt, số 1 Lê Đức Thọ",
		"weight":         1000,
		"value":          3000000,
		"transport":      "fly",
		"deliver_option": "xteam",
		"tags":           []int{1, 7},
	}
	c2 := client.NewClient(
		"https://services.ghtklab.com/services/shipment/fee?",
		"B5247cc50b15ab7C487494aD5e3F30aE41aa8bDd",
	)
	resFee, err := c2.CallGhtkApi(shippingFeeParams, http.MethodGet)
	if err != nil {
		fmt.Println(err)
	}

	var fee models.GhtkFeeDetail
	if err := mapstructure.Decode(resFee.Fee, &fee); err != nil {
		fmt.Println(err)
	}
	// feeByte, err := json.Marshal(resFee.Fee)
	// json.Unmarshal(feeByte, &fee)

	fmt.Printf("%+v\n", fee)
	fmt.Printf("%+v\n", resFee.Fee)
}

func TestGetOrderStatus(t *testing.T) {

	c3 := client.NewClient(
		"https://services.ghtklab.com/services/shipment/v2/",
		"B5247cc50b15ab7C487494aD5e3F30aE41aa8bDd",
	)
	orderLabel := "S20151392.SG8.A2.BC.300069635"
	c3.BaseURL += orderLabel
	orderStatus, err := c3.CallGhtkApi(nil, http.MethodGet)
	if err != nil {
		fmt.Println(err)
	}
	result, err := json.Marshal(orderStatus)
	fmt.Println(string(result))
	// fmt.Printf("%+v\n", orderStatus)
}

func TestPostCancelOrder(t *testing.T) {

	c4 := client.NewClient(
		"https://services.ghtklab.com/services/shipment/cancel/",
		"B5247cc50b15ab7C487494aD5e3F30aE41aa8bDd",
	)

	f, err := os.OpenFile("save-order.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()

	line, err := utils.PopLine(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	savedOrder := models.GhtkSaveOrder{}
	json.Unmarshal(line, &savedOrder)

	c4.BaseURL += savedOrder.Label
	orderCancelRes, err := c4.CallGhtkApi(nil, http.MethodPost)
	if err != nil {
		fmt.Println(err)
	}
	result, err := json.Marshal(orderCancelRes)
	fmt.Println(string(result))

	// fmt.Printf("%+v\n", orderCancelRes)
}

func TestGetHamletAddress(t *testing.T) {

	HamletAddressFeeParams := map[string]interface{}{
		"address":     "đường Đội Cấn, Ba Đình, Hà Nội",
		"province":    "Hà nội",
		"district":    "Quận Cầu Giấy",
		"ward_street": "Lê Đức Thọ",
	}
	c5 := client.NewClient(
		"https://services.ghtklab.com/services/address/getAddressLevel4?",
		"B5247cc50b15ab7C487494aD5e3F30aE41aa8bDd",
	)
	resFee, err := c5.CallGhtkApi(HamletAddressFeeParams, http.MethodGet)
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.Marshal(resFee)
	fmt.Println(string(result))
	// fmt.Printf("%+v\n", resFee)
}
