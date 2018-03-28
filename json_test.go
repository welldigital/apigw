package apigw

import (
	"errors"
	"net/http"
	"testing"
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name     string
		payload  interface{}
		expected string
		status   int
	}{
		{
			name:     "simple string",
			payload:  "simple",
			expected: `"simple"`,
			status:   http.StatusOK,
		},
		{
			name: "map",
			payload: map[string]interface{}{
				"key": "value",
			},
			expected: `{"key":"value"}`,
			status:   http.StatusLengthRequired,
		},
	}

	for _, test := range tests {
		a, err := JSON(test.payload, test.status)
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

type unmarshallable struct {
	field string
}

func (um unmarshallable) MarshalJSON() ([]byte, error) {
	return nil, errors.New("failed to marshal")
}

func TestJSONError(t *testing.T) {
	expectedStatus := http.StatusInternalServerError
	um := unmarshallable{
		field: "test",
	}
	expectedBody := `{"error":{"statusCode":500,"errorCode":"json_marshal","message":"failed to marshal JSON"}}`
	a, err := JSON(um, http.StatusInternalServerError)
	if err == nil {
		t.Fatal("expected error, but didn't get one")
	}
	if a.StatusCode != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, a.StatusCode)
	}
	if a.Body != expectedBody {
		t.Errorf("expected body '%v', got '%v'", expectedBody, a.Body)
	}
}
