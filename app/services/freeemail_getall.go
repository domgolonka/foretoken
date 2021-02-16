package services

import "github.com/domgolonka/threatdefender/app"

func FreeEmailGetDBAll(app *app.App) (*[]string, error) {
	return app.FreeEmailStore.FindAll()

}
