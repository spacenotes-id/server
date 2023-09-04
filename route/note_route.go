package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/spacenotes-id/server/controller"
	"github.com/spacenotes-id/server/middleware"
)

func RegisterNoteRoute(r fiber.Router) {
	noteController := di.
		GetInstance("noteController").(*controller.NoteController)

	r.Post("/", middleware.JwtMiddleware, noteController.CreateNote)
	r.Get("/", middleware.JwtMiddleware, noteController.FindAllNotes)
	r.Get("/trash", middleware.JwtMiddleware, noteController.FindAllTrashedNotes)
	r.Get(
		"/favorite",
		middleware.JwtMiddleware,
		noteController.FindAllFavoriteNotes,
	)
	r.Get(
		"/archive",
		middleware.JwtMiddleware,
		noteController.FindAllArchivedNotes,
	)
	r.Get(
		"/:note_id",
		middleware.JwtMiddleware, middleware.NoteOwnership,
		noteController.FindNoteByID,
	)
	r.Put(
		"/:note_id",
		middleware.JwtMiddleware, middleware.NoteOwnership,
		noteController.UpdateNote,
	)
	r.Delete(
		"/:note_id",
		middleware.JwtMiddleware, middleware.NoteOwnership,
		noteController.DeleteNote,
	)
	r.Patch(
		"/:note_id/status",
		middleware.JwtMiddleware, middleware.NoteOwnership,
		noteController.UpdateStatus,
	)
}
