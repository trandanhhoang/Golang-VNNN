package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBsonDoc(inputData interface{}) (primitive.D, error) {
	var doc primitive.D

	data, err := bson.Marshal(inputData)
	if err != nil {
		fmt.Println("Can't marshal bson", err.Error())
		return nil, err
	}
	if err = bson.Unmarshal(data, &doc); err != nil {
		fmt.Println("Can't unmarshal bson", err.Error())
		return nil, err
	}
	return doc, nil
}

func GetAndQueryBson(left interface{}, right interface{}) primitive.D {
	return bson.D{
		{"$and",
			bson.A{
				left,
				right,
			}},
	}
}
