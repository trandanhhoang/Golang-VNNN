package main

import (
	"fmt"
	"gateway/models"
	"gateway/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CreateOrder(contextGin *gin.Context) {
	data := models.OrderRequest{}
	contextGin.BindJSON(&data)

	resp, err := service.Dispatch(contextGin, &data, "/order")
	if err != nil {
		contextGin.IndentedJSON(http.StatusNotFound, gin.H{"message %s": err.Error()})
		return
	}
	jsonResponse := models.BuildBaseResponse(resp.Success)
	contextGin.JSON(http.StatusOK, jsonResponse)
}

func main() {
	//Read env if using local machine
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Error loading .env file")
	}

	//The communication protocol between user and gateway: HTTP + JSON
	port := os.Getenv("GATEWAY_PORT")
	router := gin.Default()
	router.POST("/order", CreateOrder) //

	fmt.Println("Gateway is running on port ", port)
	router.Run("localhost:" + port)
}
