package models

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func BuildBaseResponse(success bool) *BaseResponse {
	response := &BaseResponse{
		Success: false,
		Message: "Create order fail",
	}
	if success {
		response.Success = true
		response.Message = "Create order successfully"
	}
	return response
}
