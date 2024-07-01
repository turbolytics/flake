package flakehttp

import (
	"encoding/json"
	"github.com/turbolytics/flake/pkg/flake"
	"net/http"
)

type Gen interface {
	GenerateFlakeID() (flake.ID, error)
}

type Handlers struct {
	FlakeGen Gen
}

type FlakeIDResponse struct {
	ID    string   `json:"id"`
	Flake flake.ID `json:"flake"`
}

// GenerateFlakeIDHandler is an HTTP handler function to generate Flake IDs
func (h *Handlers) GenerateFlakeIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Generate a Flake ID
	id, err := h.FlakeGen.GenerateFlakeID()
	if err != nil {
		http.Error(w, "error generating flake id", http.StatusInternalServerError)
		return
	}

	p := FlakeIDResponse{
		ID:    id.String(),
		Flake: id,
	}

	// Convert the FlakeID struct to JSON
	response, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type and send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
