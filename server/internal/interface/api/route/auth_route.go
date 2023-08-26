package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/SpaceNotes/server/internal/interface/api/controller"
)

type AuthRoute struct {
	authController *controller.AuthController `di.inject:"authController"`
}

func (a *AuthRoute) RegisterRoute(r fiber.Router) {
	r.Post("/register", a.authController.Register)
	r.Post("/login", a.authController.Login)
}
