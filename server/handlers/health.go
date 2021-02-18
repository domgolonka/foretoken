package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type getHealth struct {
	HTTP bool `json:"http"`
}

func GetHealth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Response().Header.SetStatusCode(http.StatusOK)
		return c.JSON(&getHealth{
			HTTP: true,
		})
	}
}
