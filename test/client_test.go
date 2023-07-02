package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LightningDev/go-req/client"
)

func TestClientFetch(t *testing.T) {
	// Create a test server to mock the API endpoint
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the HTTP method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		// Verify the path
		if r.URL.Path != "/test" {
			t.Errorf("Expected request to /test, got %s", r.URL.Path)
		}
		// Verify the headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Write a sample response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello, world!"}`))
	}))
	defer server.Close()

	// Create a new instance of the client
	client := client.NewHttpClient(server.URL)

	// Set the required headers
	headers := &http.Header{}
	headers.Set("Content-Type", "application/json")
	client.SetHeaders(headers)

	// Send a GET request to "/test"
	response, err := client.Fetch(http.MethodGet, "/test")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// Verify the response body
	expectedResponseBody := `{"message": "Hello, world!"}`
	if string(response) != expectedResponseBody {
		t.Errorf("Expected response body:\n%s\ngot:\n%s", expectedResponseBody, string(response))
	}
}
