package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetScoreIP(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		items, err := services.ScoreIP(app, c.Params("ip"))
		if err != nil {
			return err
		}

		return c.JSON(items)
	}
}
