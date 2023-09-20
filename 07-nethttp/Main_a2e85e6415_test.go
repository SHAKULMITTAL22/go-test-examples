package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestMain_a2e85e6415 is a test function for the main function
func TestMain_a2e85e6415(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define our response
		fmt.Fprint(w, "pong")
	}))
	// Close the server when test finishes
	defer server.Close()

	// Make a request to our server with the /ping route
	resp, err := http.Get(fmt.Sprintf("%s/ping", server.URL))

	// No error should be returned
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// We expect a StatusOK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK, got %v", resp.StatusCode)
	}

	// We expect the pong response
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	if newStr != "pong" {
		t.Errorf("Expected pong as response, got %v", newStr)
	} else {
		t.Log("TestMain_a2e85e6415 passed")
	}
}
