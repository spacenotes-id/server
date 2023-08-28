package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	postgresURL  = os.Getenv("DATABASE_URL")
	PostgresPool *pgxpool.Pool
	err          error
)

func init() {
	PostgresPool, err = pgxpool.New(context.Background(), postgresURL)
	if err != nil {
		log.Fatal(err)
	}
}
