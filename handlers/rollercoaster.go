package handlers

import (
	"encoding/json"
	"fmt"
	"http-go/repositories"
	"http-go/utils"
	"net/http"
	"strconv"

	"http-go/dtos"
	"http-go/mappers"

	"github.com/go-chi/chi/v5"
)

func ListRollerCoasters(w http.ResponseWriter, r *http.Request) {
	rollerCoasters, err := repositories.GetAllRollerCoasters(r)

	if err != nil {
		http.Error(w, "Database error while querying all Roller Coasters", http.StatusInternalServerError)
		return
	}

	utils.ResponsePipe(w).JSONHeaders(http.StatusOK).JSON(rollerCoasters)
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

	utils.ResponsePipe(w).JSONHeaders(http.StatusOK).JSON(rollerCoaster)
}

func CreateRollerCoaster(w http.ResponseWriter, r *http.Request) {
	var rollerCoaster dtos.RollerCoaster

	err := json.NewDecoder(r.Body).Decode(&rollerCoaster)

	if err != nil {
		http.Error(w, "Error while Unmarshaling the Response Body", http.StatusBadRequest)
		return
	}

	rollerCoasterEnt := mappers.RollerCoasterDTOToRollerCoasterEnt(rollerCoaster)

	createdRollerCoaster, err := repositories.CreateRollerCoaster(r, rollerCoasterEnt)

	if err != nil {
		http.Error(w, "Database error while creating the Roller Coaster", http.StatusInternalServerError)
		return
	}

	utils.ResponsePipe(w).JSONHeaders(http.StatusCreated).JSON(createdRollerCoaster)
}

func UpdateRollerCoasterById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "rollerCoasterId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Roller Coaster Id", http.StatusBadRequest)
		return
	}

	var rollerCoaster dtos.RollerCoaster

	err = json.NewDecoder(r.Body).Decode(&rollerCoaster)

	if err != nil {
		http.Error(w, "Error while Unmarshaling the Response Body", http.StatusBadRequest)
		return
	}

	rollerCoasterEnt := mappers.RollerCoasterDTOToRollerCoasterEnt(rollerCoaster)

	updatedRollerCoaster, err := repositories.UpdateRollerCoasterById(r, id, rollerCoasterEnt)

	if err != nil {
		http.Error(w, fmt.Sprintf("Database error while updating the Roller Coaster with the ID %d", id), http.StatusInternalServerError)
		return
	}

	utils.ResponsePipe(w).JSONHeaders(http.StatusOK).JSON(updatedRollerCoaster)
}

func DeleteRollerCoaster(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "rollerCoasterId")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Roller Coaster Id", http.StatusBadRequest)
		return
	}

	err = repositories.DeleteRollerCoaster(r, id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Database error while deleting the Roller Coaster with the ID %d", id), http.StatusInternalServerError)
		return
	}

	utils.ResponsePipe(w).JSONHeaders(http.StatusNoContent)
}
