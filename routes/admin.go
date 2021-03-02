package routes

import (
	"github.com/gofiber/fiber/v2"
)

func admin(route fiber.Router) {
	route.Get("/", controller.admin)
	route.Get("/admin", controller.adminname)
}
