package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetEmail(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		response, err := services.EmailService(app, c.Params("email"))
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}
