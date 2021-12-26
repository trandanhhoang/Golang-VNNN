package models

type GhtkHamletAddressRequest struct {
	Province   string `json:"province"`
	District   string `json:"district"`
	WardStreet string `json:"ward_street"`
}
