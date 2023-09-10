package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	"github.com/spacenotes-id/server/config"
	"github.com/spacenotes-id/server/container"
	"github.com/spacenotes-id/server/database/postgres"
	"github.com/spacenotes-id/server/docs"
	"github.com/spacenotes-id/server/helper/exception"
	"github.com/spacenotes-id/server/route"
)

func init() {
	container.InitDI()
	docs.SwaggerInfo.Host = config.ServerHost
	docs.SwaggerInfo.Schemes = []string{config.ServerScheme}
}

func gracefullyShutdown(app *fiber.App, sigChan chan os.Signal) {
	<-sigChan
	fmt.Println("Gracefully shutting down...")

	if err := app.Shutdown(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	postgres.Pool.Close()
}

//	@title			SpaceNotes API
//	@version		1.0
//	@description	SpaceNotes REST API server

//	@contact.name	API Support
//	@contact.url	https://tfkhdyt.my.id
//	@contact.email	me@tfkhdyt.my.id

//	@license.name	MIT License
//	@license.url	https://github.com/spacenotes-id/server/blob/main/LICENSE

//	@host		localhost:8080
//	@BasePath	/v1
//	@accept		json
//	@produce	json

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				JWT key
func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			var valErr *exception.ValidationError
			if errors.As(err, &valErr) {
				errs := strings.Split(err.Error(), ";")
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"errors": errs,
				})
			}

			return c.Status(code).JSON(exception.HttpError{
				Error: err.Error(),
			})
		},
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(pprof.New())

	port := "8080"
	if config.PORT != "" {
		port = config.PORT
	}

	flag.Parse()

	v1 := app.Group("/v1")

	route.RegisterAuthRoute(v1.Group("/auth"))
	route.RegisterUserRoute(v1.Group("/users"))
	route.RegisterSpaceRoute(v1.Group("/spaces"))
	route.RegisterNoteRoute(v1.Group("/notes"))

	app.Get("/swagger/*", swagger.HandlerDefault)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go gracefullyShutdown(app, sigChan)

	if err := app.Listen(fmt.Sprintf(":%v", port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
