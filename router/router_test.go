package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Test POST /accounts
	res, err := http.Post(ts.URL+"/accounts", "application/json", nil)
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 for POST /accounts, got %v", res.StatusCode)
	}

	// Test invalid method for /accounts
	req, err := http.NewRequest("GET", ts.URL+"/accounts", nil)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected 405 for GET /accounts, got %v", res.StatusCode)
	}
}
