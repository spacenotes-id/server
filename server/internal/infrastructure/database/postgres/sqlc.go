package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tfkhdyt/SpaceNotes/server/config"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/database/postgres/sqlc"
)

func GetPostgresSQLCQuerier(ctx context.Context) *sqlc.Queries {
	conn, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		log.Fatalln("ERROR:", err.Error())
	}

	db := sqlc.New(conn)

	log.Println("INFO:", "Connected to DB")

	return db
}
