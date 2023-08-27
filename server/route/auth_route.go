package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/controller"
	"github.com/tfkhdyt/SpaceNotes/server/middleware"
)

type AuthRoute struct {
	authController *controller.AuthController `di.inject:"authController"`
}

func (a *AuthRoute) RegisterRoute(r fiber.Router) {
	r.Post("/register", a.authController.Register)
	r.Post("/login", a.authController.Login)
	r.Delete("/logout", middleware.JwtMiddleware, a.authController.Logout)
}
