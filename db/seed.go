package db

import (
	"context"
	"encoding/json"
	"http-go/dtos"
	"http-go/ent"
	"log"
	"os"
)

func Seed(client *ent.Client) {
	ctx := context.Background()

	file, err := os.ReadFile("db/seed.json")

	if err != nil {
		log.Fatalf("Could not read seed.json: %v", err)
	}

	var coasters []dtos.RollerCoaster

	if err := json.Unmarshal(file, &coasters); err != nil {
		log.Fatalf("JSON unmarshal failed: %v", err)
	}

	count, err := client.RollerCoaster.Query().Count(ctx)

	if err != nil {
		log.Fatalf("Failed counting existing records: %v", err)
	}

	if count > 0 {
		log.Println("Skipping seed: roller_coasters table already has data.")
		return
	}

	for _, c := range coasters {
		_, err := client.RollerCoaster.
			Create().
			SetName(c.Name).
			SetLocation(c.Location).
			SetHeight(c.Height).
			SetSpeed(c.Speed).
			Save(ctx)

		if err != nil {
			log.Printf("Failed to insert coaster %s: %v", c.Name, err)
		}
	}

	log.Println("Successfully seeded roller_coasters table.")
}
