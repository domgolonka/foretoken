package handlers

import (
	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetScoreIP(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ipSrv := services.IP{}
		ipSrv.Calculate(app, c.Params("ip"))
		items, err := ipSrv.ScoreIP()
		if err != nil {
			return err
		}

		return c.JSON(items)
	}
}
