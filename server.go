package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run("-mod", "vendor", "-o", "gofib")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/rohit", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", "Hello, World!")
		return c.JSON(c)
	})

	app.Listen(":3000")
}
