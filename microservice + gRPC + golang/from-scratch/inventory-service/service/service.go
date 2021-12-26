package service

import (
	"context"
	"encoding/json"
	"errors"
	db_config "inventory-service/db"
	"inventory-service/models"
	"inventory-service/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	TIME_TO_QUERRY int = 15
)

type InventoryService struct{}

func (s *InventoryService) CheckProductInInventory(ctx context.Context, data *ProductResquest) (*BaseRespond, error) {
	ctxQuerry, cancel := context.WithTimeout(context.Background(), time.Duration(TIME_TO_QUERRY)*time.Second)
	defer cancel()

	client := db_config.SetupMongoDbConnection()
	defer client.Disconnect(ctx)

	collection := client.Database("inventory").Collection("products")

	result := BaseRespond{
		Success: false,
	}

	// collection.BulkWrite()

	var productsList []models.Products
	if err := json.Unmarshal(data.GetProducts(), &productsList); err != nil {
		log.Println("Can't unmarshal products ", err.Error())
		return nil, err
	}

	for _, ele := range productsList {
		doc, err := utils.GetBsonDoc(ele.Data)
		if err != nil {
			return nil, err
		}
		nameProductBson := bson.D{{"name", ele.Name}}
		filter := utils.GetAndQueryBson(nameProductBson, doc)
		var item models.Products
		if err := collection.FindOne(ctxQuerry, filter).Decode(&item); err == mongo.ErrNoDocuments {
			return nil, errors.New("record does not exist when check")
		} else if err != nil {
			return nil, err
		}
		if item.Quantity < ele.Quantity {
			return nil, errors.New("Don't have enough item")
		}
	}

	result.Success = true
	return &result, nil
}

func (s *InventoryService) DecreaseInventory(ctx context.Context, data *ProductResquest) (*BaseRespond, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctxQuerry, cancel := context.WithTimeout(context.Background(), time.Duration(TIME_TO_QUERRY)*time.Second)
	defer cancel()

	client := db_config.SetupMongoDbConnection()
	defer client.Disconnect(ctx)

	collection := client.Database("inventory").Collection("products")

	result := BaseRespond{
		Success: true,
	}

	var productsList []models.Products
	if err := json.Unmarshal(data.GetProducts(), &productsList); err != nil {
		log.Println("Can't unmarshal products ", err.Error())
		return nil, err
	}

	for _, ele := range productsList {
		doc, err := utils.GetBsonDoc(ele.Data)
		if err != nil {
			return nil, err
		}
		nameProductBson := bson.D{{"name", ele.Name}}
		filter := utils.GetAndQueryBson(nameProductBson, doc)
		update := bson.D{{"$inc", bson.D{{"quantity", -ele.Quantity}}}}
		if _, err := collection.UpdateOne(ctxQuerry, filter, update); err != nil {
			return nil, err
		}
	}

	return &result, nil
}
