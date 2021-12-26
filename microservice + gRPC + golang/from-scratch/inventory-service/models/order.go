package models

type Tabler interface {
	TableName() string
}

type Products struct {
	Name     string      `json:"name"`
	Cost     int         `json:"cost",omitempty`
	Weight   float64     `json:"weight",omitempty`
	Quantity int         `json:"quantity",omitempty`
	Data     interface{} `json:"data",omitempty`
}
