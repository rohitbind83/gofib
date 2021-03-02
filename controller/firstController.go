package controller

import (
	"github.com/gofiber/fiber/v2"
)

// get all todos
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Rohit(c *fiber.Ctx) error {
	return c.SendString("Hello, Rohit!")
}

func admin(c *fiber.Ctx) error {
	return c.SendString("Hello, admin!")
}

func adminname(c *fiber.Ctx) error {
	return c.SendString("Hello, admin rohti!")
}
