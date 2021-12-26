package models

type GhtkOrderRequest struct {
	Products []GhtkProduct    `json:"products"`
	Order    *GhtkOrderDetail `json:"order"`
}

type GhtkOrderDetail struct {
	ID                   string        `json:"id",omitempty`
	PickName             string        `json:"pick_name",omitempty`
	PickAddress          string        `json:"pick_address",omitempty`
	PickProvince         string        `json:"pick_province",omitempty`
	PickDistrict         string        `json:"pick_district",omitempty`
	PickTel              string        `json:"pick_tel",omitempty`
	Tel                  string        `json:"tel",omitempty`
	Name                 string        `json:"name",omitempty`
	Province             string        `json:"province",omitempty`
	District             string        `json:"district",omitempty`
	Ward                 string        `json:"ward",omitempty`
	Hamlet               string        `json:"hamlet",omitempty`
	Transport            string        `json:"transport",omitempty`
	PickOption           string        `json:"pick_option",omitempty`
	DeliverOption        string        `json:"deliver_option",omitempty`
	PickSession          int           `json:"pick_session",omitempty`
	Label                string        `json:"label",omitempty`
	Area                 int           `json:"area",omitempty`
	Fee                  int           `json:"fee",omitempty`
	InsuranceFee         int           `json:"insurance_fee",omitempty`
	EstimatedPickTime    string        `json:"estimated_pick_time",omitempty`
	EstimatedDeliverTime string        `json:"estimated_deliver_time",omitempty`
	StatusID             int           `json:"status_id",omitempty`
	TrackingID           int           `json:"tracking_id",omitempty`
	SortingCode          string        `json:"sorting_code",omitempty`
	IsXfast              int           `json:"is_xfast",omitempty`
	LabelID              string        `json:"label_id",omitempty`
	PartnerID            string        `json:"partner_id",omitempty`
	Status               int           `json:"status",omitempty`
	StatusText           string        `json:"status_text",omitempty`
	Created              string        `json:"created",omitempty`
	Modified             string        `json:"modified",omitempty`
	Message              string        `json:"message",omitempty`
	PickDate             string        `json:"pick_date",omitempty`
	DeliverDate          string        `json:"deliver_date",omitempty`
	CustomerFullname     string        `json:"customer_fullname",omitempty`
	CustomerTel          string        `json:"customer_tel",omitempty`
	Address              string        `json:"address",omitempty`
	StorageDay           int           `json:"storage_day",omitempty`
	ShipMoney            int           `json:"ship_money",omitempty`
	Insurance            int           `json:"insurance",omitempty`
	Value                int           `json:"value",omitempty`
	Weight               int           `json:"weight",omitempty`
	PickMoney            int           `json:"pick_money",omitempty`
	IsFreeship           int           `json:"is_freeship",omitempty`
	Products             []GhtkProduct `json:"products",omitempty`
}

type GhtkProduct struct {
	Name        string      `json:"name",omitempty`
	FullName    string      `json:"full_name",omitempty`
	Cost        int         `json:"cost",omitempty`
	Weight      float64     `json:"weight",omitempty`
	Quantity    int         `json:"quantity",omitempty`
	ProductCode interface{} `json:"product_code",omitempty`
}
