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
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spacenotes-id/SpaceNotes/server/container"
	"github.com/spacenotes-id/SpaceNotes/server/database/postgres"
	"github.com/spacenotes-id/SpaceNotes/server/helper/exception"
	"github.com/spacenotes-id/SpaceNotes/server/route"
)

func init() {
	container.InitDI()
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

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	app.Use(recover.New())
	app.Use(pprof.New())

	port := flag.Uint("port", 8080, "server port")
	flag.Parse()

	v1 := app.Group("/v1")

	route.RegisterAuthRoute(v1.Group("/auth"))
	route.RegisterUserRoute(v1.Group("/users"))
	route.RegisterSpaceRoute(v1.Group("/spaces"))
	route.RegisterNoteRoute(v1.Group("/notes"))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go gracefullyShutdown(app, sigChan)

	if err := app.Listen(fmt.Sprintf(":%d", *port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
