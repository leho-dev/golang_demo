package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lehodev/golang-server/src/api/controllers"
)

func TodoRoute(route fiber.Router) {
	route.Get("/", controllers.GetAll)
	route.Get("/:id", controllers.GetById)
	route.Post("/", controllers.Create)
}
