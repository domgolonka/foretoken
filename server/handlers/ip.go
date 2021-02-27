package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetIP(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ipSrv := services.IP{}
		ipSrv.Calculate(app, c.Params("ip"))

		response, err := ipSrv.IPService()
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}
