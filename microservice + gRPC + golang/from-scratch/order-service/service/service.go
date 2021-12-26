package service

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"log"
	db_config "order-service/db"
	"order-service/models"
	"os"
	"time"

	grpc "google.golang.org/grpc"
	"gorm.io/gorm"
)

type OrderService struct{}

func (s *OrderService) SaveOrder(ctx context.Context, orderRequest *SaveOrderResquest) (*BaseRespond, error) {

	var order models.Order
	resp := &BaseRespond{Success: true}
	if err := json.Unmarshal(orderRequest.GetOrder(), &order); err != nil {
		fmt.Println("Can't unmarshal order ", err.Error())
	}

	var db *gorm.DB = db_config.SetupDatabaseConnectionGORM()
	defer db_config.CloseConnectionGORM(db)

	if err := db.Transaction(func(tx *gorm.DB) error {
		order.CreatedAt = time.Now()
		result := tx.Table("orders").Create(&order)
		if result.Error != nil {
			return result.Error
		}
		for _, ele := range order.Products {
			ele.OrderID = order.ID
			if err := tx.Table("products").Create(&ele).Error; err != nil {
				return err
			}
		}
		decreaseInventoryResponse, err := DecreaseInventoryService(ctx, order.Products)
		if err != nil {
			return err
		}
		resp.Success = decreaseInventoryResponse.Success
		return nil
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func DecreaseInventoryService(ctx context.Context, data interface{}) (*BaseRespond, error) {
	inventoryServicePort := os.Getenv("INVENTORY_SERVICE_PORT")
	//The server 1
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:"+inventoryServicePort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("could not start grpc client: %v", err)
		return nil, err
	}
	log.Printf("Connect with inventory service port: %v", inventoryServicePort)
	defer conn.Close()

	client := NewInventoryServiceClient(conn)

	productsRequestByte, err := json.Marshal(data)
	if err != nil {
		log.Println("Can not marshal products request by inventory service in gateway.go", err)
		return nil, err
	}

	productsRequest := ProductResquest{
		Products: productsRequestByte,
	}

	return client.DecreaseInventory(context.Background(), &productsRequest)
}
