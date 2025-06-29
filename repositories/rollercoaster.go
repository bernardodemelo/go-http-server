package repositories

import (
	"http-go/db"
	"http-go/ent"
	"http-go/ent/rollercoaster"
	"net/http"
)

func GetAllRollerCoasters(r *http.Request) ([]*ent.RollerCoaster, error) {
	rollerCoasters, err := db.Client.RollerCoaster.Query().All(r.Context())
	if err != nil {
		return nil, err
	}
	return rollerCoasters, nil
}

func GetRollerCoasterById(r *http.Request, id int) (*ent.RollerCoaster, error) {
	rollerCoaster, err := db.Client.RollerCoaster.Query().Where(rollercoaster.ID(id)).Only(r.Context())

	if err != nil {

	}
}
