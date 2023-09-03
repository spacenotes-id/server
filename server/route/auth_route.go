package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/spacenotes-id/SpaceNotes/server/controller"
	"github.com/spacenotes-id/SpaceNotes/server/middleware"
)

func RegisterAuthRoute(r fiber.Router) {
	authController := di.
		GetInstance("authController").(*controller.AuthController)

	r.Post("/register", authController.Register)
	r.Post("/login", authController.Login)
	r.Patch("/refresh", authController.Refresh)
	r.Delete("/logout", middleware.JwtMiddleware, authController.Logout)
}
