package handlers

import (
	"strings"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetDisasableEmails(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		items, err := services.DisposableGetDBAll(app)
		if err != nil {
			return err
		}
		stringByte := strings.Join(*items, "\x0A") // x20 = space and x00 = null

		return c.SendString(stringByte)
	}
}
