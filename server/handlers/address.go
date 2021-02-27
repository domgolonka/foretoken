package handlers

import (
	"github.com/domgolonka/foretoken/app/entity"
	"github.com/domgolonka/foretoken/app/services"
	"github.com/gofiber/fiber/v2"
)

func GetAddress() fiber.Handler {
	return func(c *fiber.Ctx) error {
		addy := new(entity.Address)
		if err := c.BodyParser(addy); err != nil {
			return err
		}
		response, err := services.ValidateAddress(addy)
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}
