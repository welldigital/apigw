package apigw

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// JSON returns an API Gateway response for the value v as JSON.
func JSON(v interface{}, status int) (r events.APIGatewayProxyResponse, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		r, _ = ErrorString("json_marshal", "failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	r = events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(data),
		StatusCode: status,
	}
	return
}
