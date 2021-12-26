package models

type GhtkBaseResponse struct {
	Success        bool             `json:"success"`
	Message        string           `json:"message"`
	Order          *GhtkOrderDetail `json:"order"`
	Fee            interface{}      `json:"fee"`
	Data           []string         `json:"data"`
	WarningMessage string           `json:"warning_message"`
}
