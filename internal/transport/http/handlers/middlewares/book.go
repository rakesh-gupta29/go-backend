package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LogRequests() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Request Method:", c.Method())
		return c.Next()
	}
}
