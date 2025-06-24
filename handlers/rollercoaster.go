package handlers

import (
	"encoding/json"
	"http-go/dtos"
	"net/http"
	"os"
)

type RollerCoaster struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Height   float64 `json:"height"`
	Speed    float64 `json:"speed"`
}

func RollerCoasterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the JSON
	file, err := os.Open("data/rollercoaster.json")

	// Catch the error
	if err != nil {
		http.Error(w, "Unable to load data", http.StatusInternalServerError)
		return
	}

	// Close File
	defer file.Close()

	// Initialize var coasters
	var coasters []dtos.RollerCoaster

	// Catch the error when decoding the file
	if err := json.NewDecoder(file).Decode(&coasters); err != nil {
		http.Error(w, "Error decoding data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(coasters)

}
