package model

import (
	"context"
	"fmt"
	"orderProduct/rpc/inventoryService/types"
	log "orderProduct/utils"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProductModel interface {
	UpdateMultiProduct(ctx context.Context, product []types.Product) error
}

type defaultMongoProductModel struct {
	MongoModel *mongo.Client
}

func NewMongoProductModel(url string) MongoProductModel {
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}

	return &defaultMongoProductModel{
		MongoModel: client,
	}

}

func (m *defaultMongoProductModel) UpdateMultiProduct(ctx context.Context, products []types.Product) error {
	collection := m.MongoModel.Database("inventory").Collection("products")

	var bsonID = make([]primitive.ObjectID, len(products))

	log.Info("Check products", products)

	for index, product := range products {
		var err error
		bsonID[index], err = primitive.ObjectIDFromHex(product.ID)
		if err != nil {
			return err
		}
	}

	models := make([]mongo.WriteModel, len(products))

	for index, product := range products {
		models[index] = mongo.NewUpdateOneModel().SetFilter(bson.D{{"_id", bsonID[index]}}).
			SetUpdate(bson.D{{"$inc", bson.D{{"quantity", -product.Quantity}}}})
	}

	opts := options.BulkWrite().SetOrdered(true)

	results, err := collection.BulkWrite(ctx, models, opts)
	if err != nil {
		return err
	}
	log.Info("Check mount of results ", results.UpsertedCount)
	return nil
}
