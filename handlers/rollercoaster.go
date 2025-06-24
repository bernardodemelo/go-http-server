package handlers

import (
	"encoding/json"
	"http-go/db"
	"net/http"
)

func RollerCoasterHandler(w http.ResponseWriter, r *http.Request) {

	coasters, err := db.Client.RollerCoaster.Query().All(r.Context())

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(coasters)
}
