package model

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

func getBSONObject(data interface{}) (bson.D, error) {
	var doc bson.D
	marshal, err := bson.Marshal(data)
	if err != nil {
		fmt.Println("Can't marshal bson", err.Error())
		return nil, err
	}
	if err = bson.Unmarshal(marshal, &doc); err != nil {
		fmt.Println("Can't unmarshal bson", err.Error())
		return nil, err
	}

	return doc, nil
}
