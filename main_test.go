package main

import (
	"testing"
)

// known integration test that can be made better with httptest
// will be revisited later with httptest import
func TestGet(t *testing.T) {
	resp, _, err := Get("https://www.google.com/")
	if err != nil {
		t.Fatalf("Expected no errors, go %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, recieved %d", resp.StatusCode)
	}
}
