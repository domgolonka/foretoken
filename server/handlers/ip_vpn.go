package handlers

import (
	"strings"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetVPNIPs(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		items, err := services.VpnGetDBAll(app)
		if err != nil {
			return err
		}

		stringByte := strings.Join(*items, "\x0A") // x20 = space and x00 = null

		return c.SendString(stringByte)
	}
}
