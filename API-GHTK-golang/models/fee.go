package models

type GhtkCalculateFeeRequest struct {
	PickProvince  string `json:"pick_province"`
	PickDistrict  string `json:"pick_district"`
	Province      string `json:"province"`
	District      string `json:"district"`
	Address       string `json:"address"`
	Weight        int    `json:"weight"`
	Value         int    `json:"value"`
	Transport     string `json:"transport"`
	DeliverOption string `json:"deliver_option"`
	Tags          []int  `json:"tags"`
}

type GhtkFeeDetail struct {
	Name         string                `json:"name" mapstructure:"name"`
	Fee          int                   `json:"fee" mapstructure:"fee"`
	InsuranceFee int                   `json:"insurance_fee" mapstructure:"insurance_fee"`
	IncludeVat   int                   `json:"include_vat"`
	CostId       int                   `json:"cost_id"`
	DeliveryType string                `json:"delivery_type"`
	A            string                `json:"a"`
	Dt           string                `json:"dt"`
	ExtFees      []GhtkExtendFeeDetail `json:"extFees"`
	ShipFeeOnly  int                   `json:"ship_fee_only" mapstructure:"ship_fee_only"`
	PromotionKey string                `json:"promotion_key"`
	Delivery     bool                  `json:"delivery"`
}

type GhtkExtendFeeDetail struct {
	Display string `json:"display"`
	Title   string `json:"title"`
	Amount  int    `json:"amount"`
	Type    string `json:"type"`
}
