package models

type OrderRequest struct {
	Token    string    `json:"token"`
	Products []Product `json:"products"`
}

type Product struct {
	ID       string      `json:"_id",omitempty`
	Name     string      `json:"name",omitempty`
	Weight   float64     `json:"weight",omitempty`
	Quantity int         `json:"quantity",omitempty`
	Data     interface{} `json:"data",omitempty`
}
