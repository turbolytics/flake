package flakehttp

import (
	"encoding/json"
	"github.com/turbolytics/flake/pkg/flake"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockGen implements the Gen interface for testing
type MockGen struct{}

// GenerateFlakeID generates a fixed Flake ID for testing
func (m *MockGen) GenerateFlakeID() (flake.ID, error) {
	return flake.ID{
		Timestamp: 1620000000000,
		WorkerID:  0x123456,
		Sequence:  42,
	}, nil
}

func TestGenerateFlakeIDHandler(t *testing.T) {
	// Create the Handlers instance with the mock generator
	handlers := &Handlers{
		FlakeGen: &MockGen{},
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/generate", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GenerateFlakeIDHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expectedID := flake.ID{
		Timestamp: 1620000000000,
		WorkerID:  0x123456,
		Sequence:  42,
	}
	expectedResponse := FlakeIDResponse{
		ID:    expectedID.String(),
		Flake: expectedID,
	}

	var response FlakeIDResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("unable to parse response: %v", err)
	}

	if response.ID != expectedResponse.ID {
		t.Errorf("handler returned unexpected ID: got %v want %v",
			response.ID, expectedResponse.ID)
	}

	if response.Flake != expectedResponse.Flake {
		t.Errorf("handler returned unexpected Flake ID struct: got %+v want %+v",
			response.Flake, expectedResponse.Flake)
	}
}
