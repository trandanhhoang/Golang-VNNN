package model

//go:generate goctl model mongo -t User
import "github.com/globalsign/mgo/bson"

type Product struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `json:"name"`
	Size     string        `json:"size"`
	Quantity int           `json:"quantity"`
	Length   string        `json:"length"`
}
