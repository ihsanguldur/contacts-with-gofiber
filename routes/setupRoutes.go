package routes

import "github.com/gofiber/fiber/v2"

//SetupRoutes is a function which create /api sub-route
//for user router, auth router and contact router
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	UserRouter(api)
	AuthRouter(api)
	ContactRouter(api)
}
