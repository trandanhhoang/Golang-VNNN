package models

type OrderRequest struct {
	Token    string    `json:"token"`
	Products []Product `json:"products"`
}

type Product struct {
	Name     string      `json:"name",omitempty`
	Cost     int         `json:"cost",omitempty`
	Weight   float64     `json:"weight",omitempty`
	Quantity int         `json:"quantity",omitempty`
	Data     interface{} `json:"data",omitempty`
}
