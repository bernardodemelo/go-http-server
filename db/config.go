package db

import (
	"context"
	"encoding/json"
	"fmt"
	"http-go/dtos"
	"http-go/ent"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Client *ent.Client

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (fallback to system ENV)")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port,
		user, dbname, password,
	)

	var err error

	Client, err = ent.Open("postgres", dsn)

	if Client == nil {
		log.Fatalf("Failed to create DB Client")
	}

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Schema migration failed: %v", err)
	}

}

func Seed() {
	ctx := context.Background()

	count, err := Client.RollerCoaster.Query().Count(ctx)

	if count > 0 {
		log.Println("Skipping seed: roller_coasters table already has data.")
		return
	}

	if err != nil {
		log.Fatalf("Failed counting existing records: %v", err)
	}

	file, err := os.ReadFile("db/seed.json")

	if err != nil {
		log.Fatalf("Could not read seed.json: %v", err)
	}

	var coasters []dtos.RollerCoaster

	if err := json.Unmarshal(file, &coasters); err != nil {
		log.Fatalf("JSON unmarshal failed: %v", err)
	}

	for _, c := range coasters {
		_, err := Client.RollerCoaster.
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
