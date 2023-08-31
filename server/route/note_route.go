package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/controller"
	"github.com/tfkhdyt/SpaceNotes/server/middleware"
)

func RegisterNoteRoute(r fiber.Router) {
	noteController := di.
		GetInstance("noteController").(*controller.NoteController)

	r.Post(
		"/",
		middleware.JwtMiddleware,
		noteController.CreateNote,
	)
}
