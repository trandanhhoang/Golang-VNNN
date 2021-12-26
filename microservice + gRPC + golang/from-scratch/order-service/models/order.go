package models

import (
	"time"

	"gorm.io/datatypes"
)

type Tabler interface {
	TableName() string
}

type Order struct {
	ID        int       `json:"id" `
	Token     string    `json:"token"`
	Products  []Product `json:"products"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	OrderID  int            `json:"order_id"`
	Name     string         `json:"name"`
	Cost     int            `json:"cost",omitempty`
	Weight   float64        `json:"weight",omitempty`
	Quantity int            `json:"quantity",omitempty`
	Data     datatypes.JSON `json:"data",omitempty`
}
