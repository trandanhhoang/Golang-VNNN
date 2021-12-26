package service

import (
	"encoding/json"
	"gateway/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func SaveOrderToDBByGRPC(contextGin *gin.Context, order *models.OrderRequest) SameReturnFunc {
	orderServicePort := os.Getenv("ORDER_SERVICE_PORT")

	orderRequestByte, err := json.Marshal(order)
	if err != nil {
		log.Println("Can not marshal products request by inventory service in gateway.go", err)
	}

	orderRequest := SaveOrderResquest{
		Order: orderRequestByte,
	}

	return func() (*BaseRespond, error) { return SaveOrder(orderServicePort, &orderRequest) }
}

func CheckProductIsSuffcientByGRPC(contextGin *gin.Context, order *models.OrderRequest) SameReturnFunc {
	inventoryServicePort := os.Getenv("INVENTORY_SERVICE_PORT")

	productsRequestByte, err := json.Marshal(order.Products)
	if err != nil {
		log.Println("Can not marshal products request by inventory service in gateway.go", err)
	}
	productsRequest := ProductResquest{
		Products: productsRequestByte,
	}

	return func() (*BaseRespond, error) { return CheckProductInInventory(inventoryServicePort, &productsRequest) }
}

func CreateOrderByGRPC(contextGin *gin.Context, token string) SameReturnFunc {
	userServicePort := os.Getenv("USER_SERVICE_PORT")

	tokenRequest := TokenRequest{
		Token: token,
	}

	return func() (*BaseRespond, error) { return VerifyUserByToken(userServicePort, &tokenRequest) }
}
