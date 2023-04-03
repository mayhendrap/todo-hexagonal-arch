package dto

type BaseResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data"`
}

func ToResponse(message string, status int, data any) BaseResponse {
	return BaseResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
