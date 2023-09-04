package config

import "os"

var (
	ServerHost   = os.Getenv("SERVER_HOST")
	ServerScheme = os.Getenv("SERVER_SCHEME")
)
