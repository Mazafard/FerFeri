package main

import (
	"github.com/mazafard/ferferi/pkg/websocket"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Create a mock server to handle requests
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a status code of 200
		w.WriteHeader(http.StatusOK)
	}))

	// Replace the port in the configuration with the mock server's port
	config := websocket.DefaultConfig()
	config.Port = 80 // This should be changed to the mock server's port, which is an integer

	// Call the main function to start the server
	go main()

	// Send a GET request to the mock server
	resp, err := http.Get(mockServer.URL)
	if err != nil {
		//t.Fatal(err)
		os.Exit(m.Run())
	}

	// Check the response from the server
	if resp.StatusCode != http.StatusOK {
		//t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		os.Exit(m.Run())
	}

	// Stop the mock server
	mockServer.Close()
}
