package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	// Render index template
	return c.Render("info", nil)
}
