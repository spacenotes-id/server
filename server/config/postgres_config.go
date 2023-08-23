package config

import "os"

var PostgresURL = os.Getenv("DATABASE_URL")
