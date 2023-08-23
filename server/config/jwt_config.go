package config

import "os"

var JwtAccessTokenKey = os.Getenv("JWT_ACCESS_TOKEN_KEY")
