package postgres

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/tfkhdyt/SpaceNotes/server/config"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
)

func GetPostgresSQLCQuerier() (*sqlc.Queries, error) {
	db := sqlc.New(config.PostgresPool)

	log.Info("Connected to DB")

	return db, nil
}
