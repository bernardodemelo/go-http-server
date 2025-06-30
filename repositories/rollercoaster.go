package repositories

import (
	"fmt"
	"http-go/db"
	"http-go/ent"
	"http-go/ent/rollercoaster"
	"net/http"
)

func GetAllRollerCoasters(r *http.Request) ([]*ent.RollerCoaster, error) {
	rollerCoasters, err := db.Client.RollerCoaster.Query().All(r.Context())

	if err != nil {
		return nil, fmt.Errorf("failed to get the list of roller coasters: %w", err)
	}

	return rollerCoasters, nil
}

func GetRollerCoasterById(r *http.Request, id int) (*ent.RollerCoaster, error) {
	rollerCoaster, err := db.Client.RollerCoaster.Query().Where(rollercoaster.ID(id)).Only(r.Context())

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("roller Coaster with the Id %d not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to get roller coaster with the Id %d: %w", id, err)
	}

	return rollerCoaster, nil
}
