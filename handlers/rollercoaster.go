package handlers

import (
	"encoding/json"
	"http-go/db"
	"http-go/ent"
	"http-go/ent/rollercoaster"
	"http-go/repositories"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ListRollerCoasters(w http.ResponseWriter, r *http.Request) {
	rollerCoasters, err := repositories.GetAllRollerCoasters(r)

	if err != nil {
		http.Error(w, "Database error while querying all Roller Coasters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(rollerCoasters)
}

func GetRollerCoaster(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "rollerCoasterId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Roller Coaster Id", http.StatusBadRequest)
		return
	}

	coaster, err := db.Client.RollerCoaster.Query().Where(rollercoaster.ID(id)).Only(r.Context())

	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "Roller coaster not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(coaster)

}
