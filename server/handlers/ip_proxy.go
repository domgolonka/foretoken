package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetProxyIPs(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		items, err := services.ProxyGetDBAll(app)
		if err != nil {
			return err
		}

		return c.JSON(items)
	}
}
