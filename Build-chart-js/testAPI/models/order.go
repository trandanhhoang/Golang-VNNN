package models

type OrderRequest struct {
	Type string `json:"type,omitempty"`
	Data Data   `json:"data"`
	Options interface{}  `json:"options,omitempty"`
}

type Data struct {
	Labels   interface{} `json:"labels"`
	Datasets interface{}  `json:"datasets"`
}

// type Datasets struct {
	// Label           string      `json:"label"`
	// Data            interface{} `json:"data"`
	// BackgroundColor interface{} `json:"backgroundColor,omitempty"`
	// BorderColor     interface{} `json:"borderColor,omitempty"`
	// BorderWidth     int         `json:"borderWidth,omitempty"`
// }
