package routes

import (
	"github.com/gofiber/fiber/v2"
	"rehber/handlers"
	"rehber/middlewares"
)

//ContactRouter is a function which create /contacts sub-route
func ContactRouter(app fiber.Router) {
	api := app.Group("/contacts", middlewares.Protected())

	api.Post("/:id", handlers.CreateContact)
}
