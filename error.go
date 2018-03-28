package apigw

import (
	"github.com/aws/aws-lambda-go/events"
)

// ErrorResult is for creating a JSON response of `{ "statusCode": 400, "errorCode": "reason_failed", "message": "didn't work" }`.
type ErrorResult struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

// ErrorResponse matches the standard:
// {
// 		"error": {
// 		"statusCode": 404,
// 		"errorCode": "UserNotFound",
// 		"message": "The account could not be found"
// 	}
// }
//
type ErrorResponse struct {
	Error ErrorResult `json:"error"`
}

// NewErrorResponse creates a standard error response.
func NewErrorResponse(statusCode int, errorCode string, message string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorResult{
			StatusCode: statusCode,
			ErrorCode:  errorCode,
			Message:    message,
		},
	}
}

// Error returns a JSON ErrorResult.
func Error(code string, err error, status int) (events.APIGatewayProxyResponse, error) {
	response := NewErrorResponse(status, code, err.Error())
	return JSON(response, status)
}

// ErrorString returns a JSON ErrorResult.
func ErrorString(code string, message string, status int) (events.APIGatewayProxyResponse, error) {
	response := NewErrorResponse(status, code, message)
	return JSON(response, status)
}
