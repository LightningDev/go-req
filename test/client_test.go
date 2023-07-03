package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	client "github.com/LightningDev/go-req"
)

func TestClientFetch(t *testing.T) {
	// Create a test server to mock the API endpoint
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.URL.Path != "/test" {
			t.Errorf("Expected request to /test, got %s", r.URL.Path)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello, world!"}`))
	}))
	defer server.Close()

	client := client.New(server.URL)

	headers := &http.Header{}
	headers.Set("Content-Type", "application/json")
	client.SetHeadersFromObject(headers)

	response, err := client.Fetch(http.MethodGet, "/test")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	expectedResponseBody := `{"message": "Hello, world!"}`
	if string(response) != expectedResponseBody {
		t.Errorf("Expected response body:\n%s\ngot:\n%s", expectedResponseBody, string(response))
	}
}
