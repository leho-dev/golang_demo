package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ProtectedMiddleware(c *fiber.Ctx) error {
	log.Println("---> ProtectedMiddleware through! <---")
	c.Next()
	return nil
}
