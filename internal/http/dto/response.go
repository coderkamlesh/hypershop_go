package dto

import "github.com/coderkamlesh/hypershop_go/internal/constants"

// GlobalResponse is the standard response structure for all endpoints
type GlobalResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success returns a success response with data (status = 1)
func Success(message string, data interface{}) GlobalResponse {
	return GlobalResponse{
		Status:  constants.SUCCESS,
		Message: message,
		Data:    data,
	}
}

// SuccessWithoutData returns a success response without data
func SuccessWithoutData(message string) GlobalResponse {
	return GlobalResponse{
		Status:  constants.SUCCESS,
		Message: message,
		Data:    nil,
	}
}

// Failure returns a failure response (status = 0)
func Failure(message string) GlobalResponse {
	return GlobalResponse{
		Status:  constants.FAILED,
		Message: message,
		Data:    nil,
	}
}

// InvalidToken returns invalid token response (status = 3)
func InvalidToken() GlobalResponse {
	return GlobalResponse{
		Status:  constants.INVALID_TOKEN,
		Message: constants.InvalidTokenMessage,
		Data:    nil,
	}
}

// InvalidTokenWithMessage returns invalid token with custom message
func InvalidTokenWithMessage(message string) GlobalResponse {
	return GlobalResponse{
		Status:  constants.INVALID_TOKEN,
		Message: message,
		Data:    nil,
	}
}
