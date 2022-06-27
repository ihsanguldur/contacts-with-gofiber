package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rehber/database"
	"rehber/routes"
	"rehber/utils"
)

func main() {
	database.Connect()

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
