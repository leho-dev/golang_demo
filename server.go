package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/lehodev/golang-server/src/api/middlewares"
	"github.com/lehodev/golang-server/src/api/routes"
	"github.com/lehodev/golang-server/src/configs/db/mongodb"
)

func main() {
	mongodb.Connect()
	defer func() {
		if err := mongodb.Disconnect(); err != nil {
			log.Println("---> Error disconnecting from MongoDB:", err)
		}
	}()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(middlewares.ProtectedMiddleware)

	api := app.Group("/api")

	routes.TodoRoute(api)

	app.Listen("localhost:8080")
}
