package db

import (
	"context"
	"fmt"
	"http-go/ent"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *ent.Client {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (fallback to system ENV)")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
		user, dbname, password,
	)

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Schema migration failed: %v", err)
	}

	return client
}
