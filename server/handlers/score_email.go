package handlers

import (
	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetScoreEmail(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		emailSrv := services.Email{}
		emailSrv.Calculate(app, c.Params("email"))
		items, err := emailSrv.ScoreEmail()
		if err != nil {
			return err
		}

		return c.JSON(items)
	}
}
