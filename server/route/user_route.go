package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/controller"
	"github.com/tfkhdyt/SpaceNotes/server/middleware"
)

type UserRoute struct {
	userController *controller.UserController `di.inject:"userController"`
}

func (u *UserRoute) RegisterRoute(r fiber.Router) {
	r.Get("/me", middleware.JwtMiddleware, u.userController.FindMyAccount)
	r.Put("/me", middleware.JwtMiddleware, u.userController.UpdateMyAccount)
	r.Patch("me/email", middleware.JwtMiddleware, u.userController.UpdateMyEmail)
	r.Patch(
		"me/password",
		middleware.JwtMiddleware,
		u.userController.UpdateMyPassword,
	)
	r.Delete("me", middleware.JwtMiddleware, u.userController.DeleteMyAccount)
}
