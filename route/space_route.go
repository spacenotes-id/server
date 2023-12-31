package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/spacenotes-id/server/controller"
	"github.com/spacenotes-id/server/middleware"
)

func RegisterSpaceRoute(r fiber.Router) {
	spaceController := di.
		GetInstance("spaceController").(*controller.SpaceController)
	noteController := di.
		GetInstance("noteController").(*controller.NoteController)

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
	r.Delete(
		"/:space_id",
		middleware.JwtMiddleware, middleware.SpaceOwnership,
		spaceController.DeleteSpace,
	)

	r.Get(
		"/:space_id/notes",
		middleware.JwtMiddleware, middleware.SpaceOwnership,
		noteController.FindAllNotesBySpaceID,
	)
}
