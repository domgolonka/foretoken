package handlers

import (
	"github.com/domgolonka/foretoken/app"
	"github.com/gofiber/fiber/v2"
)

// todo
func GetDCIPs(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.SendString("")
	}
}
