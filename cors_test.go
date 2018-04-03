package apigw

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestSetCORSHeaders(t *testing.T) {
	tests := []struct {
		name             string
		r                events.APIGatewayProxyResponse
		allowOrigin      string
		allowCredentials bool
		expected         map[string]string
	}{
		{
			name: "basic",
			r: events.APIGatewayProxyResponse{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body:       "OK",
				StatusCode: 200,
			},
			allowOrigin: "*",
			expected: map[string]string{
				"Content-Type":                     "application/json",
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "false",
			},
		},
		{
			name: "nil map",
			r: events.APIGatewayProxyResponse{
				Body:       "OK",
				StatusCode: 200,
			},
			allowOrigin: "*",
			expected: map[string]string{
				"Access-Control-Allow-Origin":      "*",
				"Access-Control-Allow-Credentials": "false",
			},
		},
		{
			name: "allow origin",
			r: events.APIGatewayProxyResponse{
				Body:       "OK",
				StatusCode: 200,
			},
			allowOrigin: "test",
			expected: map[string]string{
				"Access-Control-Allow-Origin":      "test",
				"Access-Control-Allow-Credentials": "false",
			},
		},
		{
			name: "allow credentials",
			r: events.APIGatewayProxyResponse{
				Body:       "OK",
				StatusCode: 200,
			},
			allowOrigin:      "test",
			allowCredentials: true,
			expected: map[string]string{
				"Access-Control-Allow-Origin":      "test",
				"Access-Control-Allow-Credentials": "true",
			},
		},
	}

	for _, test := range tests {
		SetCORSHeaders(test.allowOrigin, test.allowCredentials, &test.r)
		if !reflect.DeepEqual(test.r.Headers, test.expected) {
			t.Errorf("%s: expected '%v', got '%v'", test.name, test.expected, test.r.Headers)
		}
	}
}

func TestSetCORSHeadersDoesNotPanic(t *testing.T) {
	SetCORSHeaders("*", true, nil)
}
