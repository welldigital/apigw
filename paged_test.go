package apigw

import (
	"net/http"
	"testing"
)

func TestPaged(t *testing.T) {
	tests := []struct {
		name       string
		startIndex int
		pageSize   int
		nextIndex  int
		payload    []interface{}
		expected   string
		status     int
	}{
		{
			name:     "simple string",
			payload:  []interface{}{"simple"},
			expected: `{"startIndex":0,"pageSize":0,"data":["simple"]}`,
			status:   http.StatusOK,
		},
		{
			name: "map",
			payload: []interface{}{map[string]interface{}{
				"key": "value",
			},
			},
			startIndex: 5,
			expected:   `{"startIndex":5,"pageSize":0,"data":[{"key":"value"}]}`,
			status:     http.StatusLengthRequired,
		},
		{
			name: "mixed array",
			payload: []interface{}{
				123,
				"test",
			},
			pageSize: 50,
			expected: `{"startIndex":0,"pageSize":50,"data":[123,"test"]}`,
			status:   http.StatusLengthRequired,
		},
		{
			name:      "next index",
			payload:   []interface{}{"simple"},
			nextIndex: 1,
			expected:  `{"startIndex":0,"pageSize":0,"nextIndex":1,"data":["simple"]}`,
			status:    http.StatusOK,
		},
	}

	for _, test := range tests {
		a, err := Paged(test.startIndex, test.pageSize, test.nextIndex, test.payload, test.status)
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
