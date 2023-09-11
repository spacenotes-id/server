package config

import (
	"os"
)

var (
	PostgresURL  = os.Getenv("DATABASE_URL")
	MigrationDir = os.Getenv("DBMATE_MIGRATIONS_DIR")
	SchemaFile   = os.Getenv("DBMATE_SCHEMA_FILE")
)
