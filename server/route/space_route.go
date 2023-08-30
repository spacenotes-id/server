package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/controller"
	"github.com/tfkhdyt/SpaceNotes/server/middleware"
)

func RegisterSpaceRoute(r fiber.Router) {
	spaceController := di.
		GetInstance("spaceController").(*controller.SpaceController)

	r.Post("/", middleware.JwtMiddleware, spaceController.CreateSpace)
	r.Get("/", middleware.JwtMiddleware, spaceController.FindAllSpacesByUserID)
	r.Get(
		"/:space_id",
		middleware.JwtMiddleware, middleware.SpaceOwnership,
		spaceController.FindSpaceByID,
	)
	r.Put(
		"/:space_id",
		middleware.JwtMiddleware, middleware.SpaceOwnership,
		spaceController.UpdateSpace,
	)
}
