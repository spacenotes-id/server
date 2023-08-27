package postgres

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tfkhdyt/SpaceNotes/server/config"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
)

func GetPostgresSQLCQuerier(ctx context.Context) (*sqlc.Queries, error) {
	conn, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		return nil, err
	}

	db := sqlc.New(conn)

	log.Info("Connected to DB")

	return db, nil
}
