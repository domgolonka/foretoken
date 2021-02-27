package services

import "github.com/domgolonka/foretoken/app"

func FreeEmailGetDBAll(app *app.App) (*[]string, error) {
	return app.FreeEmailStore.FindAll()

}
