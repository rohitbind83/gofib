package routes

import (
	"learn/controller"

	"github.com/gofiber/fiber/v2"
)

func users(route fiber.Router) {
	route.Get("/", controller.Hello)
	route.Get("/rohit", controller.Rohit)
}
