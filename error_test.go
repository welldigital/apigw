package apigw

import (
	"errors"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		err      error
		expected string
		status   int
	}{
		{
			name:     "basic",
			code:     "basic",
			err:      errors.New("basic"),
			expected: `{"error":{"statusCode":500,"errorCode":"basic","message":"basic"}}`,
			status:   http.StatusInternalServerError,
		},
		{
			name:     "unathorized",
			code:     "unathorized",
			err:      errors.New("user not authorized"),
			expected: `{"error":{"statusCode":401,"errorCode":"unathorized","message":"user not authorized"}}`,
			status:   http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		a, err := Error(test.code, test.err, test.status)
		if err != nil {
			t.Fatalf("%s: failed to create response: %v", test.name, err)
		}
		if a.StatusCode != test.status {
			t.Errorf("%s: expected status %d, got %d", test.name, test.status, a.StatusCode)
		}
		if a.Body != test.expected {
			t.Errorf("%s: expected body '%v', got '%v'", test.name, test.expected, a.Body)
		}
	}
}

func TestErrorString(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		err      string
		expected string
		status   int
	}{
		{
			name:     "basic",
			code:     "basic",
			err:      "basic message",
			expected: `{"error":{"statusCode":500,"errorCode":"basic","message":"basic message"}}`,
			status:   http.StatusInternalServerError,
		},
		{
			name:     "unathorized",
			code:     "unathorized",
			err:      "user not authorized message",
			expected: `{"error":{"statusCode":401,"errorCode":"unathorized","message":"user not authorized message"}}`,
			status:   http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		a, err := ErrorString(test.code, test.err, test.status)
		if err != nil {
			t.Fatalf("%s: failed to create response: %v", test.name, err)
		}
		if a.StatusCode != test.status {
			t.Errorf("%s: expected status %d, got %d", test.name, test.status, a.StatusCode)
		}
		if a.Body != test.expected {
			t.Errorf("%s: expected body '%v', got '%v'", test.name, test.expected, a.Body)
		}
	}
}
