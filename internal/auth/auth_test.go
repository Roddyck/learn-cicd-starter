package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headers  http.Header
		expected string
		err      error
	}{
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			expected: "1234567890",
			err:      nil,
		},
		{
			headers: http.Header{
				"Authorization": []string{},
			},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
	}

	for _, test := range tests {
		actual, err := GetAPIKey(test.headers)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("expected error %v, got %v", test.err, err)
		}
		if actual != test.expected {
			t.Errorf("expected %v, got %v", test.expected, actual)
		}
	}
}
