package http

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/internal/interface/api/route"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
)

func StartFiberServer() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *exception.HttpError
			if errors.As(err, &e) {
				code = int(e.StatusCode())
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

	di.GetInstance("userRoute").(*route.UserRoute).
		RegisterRoute(v1.Group("/users"))
	di.GetInstance("authRoute").(*route.AuthRoute).
		RegisterRoute(v1.Group("/auth"))

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", *port)))
}
