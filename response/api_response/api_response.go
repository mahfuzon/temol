package api_response

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ConverseToSuccessResponse(message string, data interface{}) ApiResponse {
	successResponse := ApiResponse{
		Status:  "ok",
		Message: message,
		Data:    data,
	}

	return successResponse
}

func ConverseToErrorResponse(message string, data interface{}) ApiResponse {
	errorResponse := ApiResponse{
		Status:  "error",
		Message: message,
		Data:    data,
	}

	return errorResponse
}
