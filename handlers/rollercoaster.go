package handlers

import (
	"fmt"
	"http-go/repositories"
	"http-go/utils"
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

	utils.WriteInJSON(w, http.StatusOK, rollerCoasters)
}

func GetRollerCoaster(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "rollerCoasterId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Roller Coaster Id", http.StatusBadRequest)
		return
	}

	rollerCoaster, err := repositories.GetRollerCoasterById(r, id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Database error while querying the Roller Coaster with the ID %d", id), http.StatusInternalServerError)
		return
	}

	utils.WriteInJSON(w, http.StatusOK, rollerCoaster)

}
