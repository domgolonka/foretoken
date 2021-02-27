package server

import (
	"github.com/domgolonka/foretoken/app"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type message struct {
	Message string `json:"message"`
}

func Error(app *app.App) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		code := fiber.StatusInternalServerError
		app.Logger.Error(err)

		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(message{
				Message: e.Message,
			})
		}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(message{Message: "Data is invalid"})
		}

		//if err, ok := err.(validator.ValidationErrors); ok {
		//	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Translate(translator))
		//}

		return ctx.Status(code).JSON(message{Message: "An error has occurred!"})
	}
}
