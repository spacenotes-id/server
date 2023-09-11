package postgres

import (
	"context"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/spacenotes-id/server/config"
	"github.com/spacenotes-id/server/database/postgres/sqlc"
)

var (
	Pool *pgxpool.Pool
	err  error
)

func init() {
	Pool, err = pgxpool.New(context.Background(), config.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}

	migrate()
}

func migrate() {
	u, errUrl := url.Parse(config.PostgresURL)
	if errUrl != nil {
		log.Fatal(errUrl)
	}

	migration := dbmate.New(u)
	migration.MigrationsDir = []string{config.MigrationDir}
	migration.SchemaFile = config.SchemaFile

	if err := migration.CreateAndMigrate(); err != nil {
		log.Fatal(err)
	}
}

func GetPostgresSQLCQuerier() (*sqlc.Queries, error) {
	db := sqlc.New(Pool)

	log.Info("Connected to DB")

	return db, nil
}
