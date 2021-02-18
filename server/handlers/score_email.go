package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetScoreEmail(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		items, err := services.ScoreEmail(app, c.Get("email"))
		if err != nil {
			return err
		}

		return c.JSON(items)
	}
}
