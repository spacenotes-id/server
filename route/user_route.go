package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"

	"github.com/spacenotes-id/server/controller"
	"github.com/spacenotes-id/server/middleware"
)

func RegisterUserRoute(r fiber.Router) {
	userController := di.
		GetInstance("userController").(*controller.UserController)

	r.Get("/me", middleware.JwtMiddleware, userController.FindMyAccount)
	r.Put("/me", middleware.JwtMiddleware, userController.UpdateMyAccount)
	r.Patch("/me/email", middleware.JwtMiddleware, userController.UpdateMyEmail)
	r.Patch(
		"/me/password",
		middleware.JwtMiddleware,
		userController.UpdateMyPassword,
	)
	r.Delete("/me", middleware.JwtMiddleware, userController.DeleteMyAccount)
}
