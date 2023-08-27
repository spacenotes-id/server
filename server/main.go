package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/container"
	"github.com/tfkhdyt/SpaceNotes/server/helper/exception"
	"github.com/tfkhdyt/SpaceNotes/server/route"
)

func init() {
	container.InitDi()
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
	// app.Use(recover.New())
	app.Use(pprof.New())

	port := flag.Uint("port", 8080, "server port")
	flag.Parse()

	v1 := app.Group("/v1")

	di.GetInstance("userRoute").(*route.UserRoute).
		RegisterRoute(v1.Group("/users"))
	di.GetInstance("authRoute").(*route.AuthRoute).
		RegisterRoute(v1.Group("/auth"))

	log.Info("Server is running at port ", *port)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
