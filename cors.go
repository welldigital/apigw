package apigw

import "github.com/aws/aws-lambda-go/events"

// SetCORSHeaders sets CORS headers on the response.
func SetCORSHeaders(allowOrigin string, allowCredentials bool, r *events.APIGatewayProxyResponse) {
	if r == nil {
		return
	}
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}
	r.Headers["Access-Control-Allow-Origin"] = allowOrigin
	if allowCredentials {
		r.Headers["Access-Control-Allow-Credentials"] = "true"
	} else {
		r.Headers["Access-Control-Allow-Credentials"] = "false"
	}
}
