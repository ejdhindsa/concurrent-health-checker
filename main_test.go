package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name           string
		testURL        string
		mockStatusCode int
		expectError    bool
	}{
		{"Valid 200 OK", "", http.StatusOK, false},
		{"Not Found 404", "", http.StatusNotFound, false},
		{"Server Error 500", "", http.StatusInternalServerError, false},
		{"Network Failure", "http://this-is-not-a-real-website-ever.com", 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.mockStatusCode)
			}))
			defer server.Close()

			urlToTest := server.URL
			if tc.testURL != "" {
				urlToTest = tc.testURL
			}

			resp, _, err := Get(urlToTest)

			if tc.expectError && err == nil {
				t.Errorf("Expected an error, but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Did not expect an error, but got: %v", err)
			}

			if resp != nil && tc.mockStatusCode != 0 && resp.StatusCode != tc.mockStatusCode {
				t.Errorf("Expected status %d, got %d", tc.mockStatusCode, resp.StatusCode)
			}
		})
	}
}
