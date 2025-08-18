package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
	 name     string
	 header  http.Header
	 expected string
	 error    error
	}{
	 {name: "key", header: http.Header{"Authorization": []string{"ApiKey 123"}}, expected: "123", error: nil},
	 {name: "empty", header: http.Header{"Authorization": []string{}}, expected: "", error: ErrNoAuthHeaderIncluded},
	}

	for _, tt := range tests {
 			t.Run(tt.name, func(t *testing.T) {
 				got, err := GetAPIKey(tt.header)
 				if (err != tt.error) {
 					t.Errorf("GetAPIKey(%q) error = %v; want %v", tt.header.Get("Authorization"), err, tt.error)
 				}
 				if got != tt.expected {
 					t.Errorf("GetAPIKey(%q) = %q; want %q", tt.header.Get("Authorization"), got, tt.expected)
 				}
 			})
	}
}