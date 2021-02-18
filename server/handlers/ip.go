package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetIP(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		response, err := services.IPService(app, c.Params("ip"))
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}
