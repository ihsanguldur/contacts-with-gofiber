package routes

import (
	"github.com/gofiber/fiber/v2"
	"rehber/handlers"
)

//AuthRouter is a function which create /auth sub-route
func AuthRouter(app fiber.Router) {
	api := app.Group("/auth")

	api.Post("/login", handlers.Login)
}
