package handlers

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
	"github.com/gofiber/fiber/v2"
)

type validEmail struct {
	Valid bool `json:"valid"`
}

func GetValidateEmail(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := services.ValidateEmail(app, c.Params("email"))
		if err != nil {
			return c.JSON(&validEmail{
				Valid: false,
			})
		}

		return c.JSON(&validEmail{
			Valid: true,
		})
	}
}
