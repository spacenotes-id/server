package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/config"
)

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey:  jwtware.SigningKey{Key: []byte(config.JwtAccessTokenKey)},
	TokenLookup: "header:Authorization,cookie:accessToken",
	AuthScheme:  "Bearer",
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	},
})
