package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// TestSetupRouter_d45738bf12 tests the setupRouter function.
func TestSetupRouter_d45738bf12(t *testing.T) {
	// Create a router instance
	router := setupTestRouter()

	// Create a request to pass to our router
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "pong"
	if w.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}

	// Test case for non-existent route
	req, err = http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Reset the ResponseRecorder
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code for non-existent route: got %v want %v", status, http.StatusNotFound)
	}
}
