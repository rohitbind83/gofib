package main

import (
	"learn/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setUpRoute(app)

	app.Listen(":3000")
}

func setUpRoute(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	// api := app.Group("/api")
	// app.Get("/rohit", func(c *fiber.Ctx) error {
	// 	return c.JSON(c)
	// })
	routes.users(app)
	routes.TodoRoute(app.Group("/admin"))
}
