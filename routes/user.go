package routes

import (
	"github.com/gofiber/fiber/v2"
	"rehber/handlers"
	"rehber/middlewares"
)

//UserRouter is a function which create /user sub-route
func UserRouter(app fiber.Router) {
	api := app.Group("/user")

	api.Post("/", handlers.CreateUser)
	api.Delete("/:id", middlewares.Protected(), handlers.DeleteUser)
	api.Put("/:id", middlewares.Protected(), handlers.UpdateUser)
	api.Get("/:id", middlewares.Protected(), handlers.SearchUser)
	api.Get("/contacts/:id", middlewares.Protected(), handlers.ListContacts)
}
