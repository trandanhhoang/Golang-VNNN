package types

type OrderReq struct {
	ID       int       `json:"id" `
	Token    string    `json:"token"`
	Products []Product `json:"products"`
}

type Product struct {
	OrderID  int               `json:"order_id"`
	ID       string            `json:"_id"`
	Name     string            `json:"name"`
	Quantity int               `json:"quantity"`
	Data     map[string]string `json:"data"`
}
