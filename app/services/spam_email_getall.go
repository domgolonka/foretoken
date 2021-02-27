package services

import (
	"github.com/domgolonka/foretoken/app"
)

func SpamEmailGetCacheAll(app *app.App) (*[]string, error) {
	return app.SpamEmailGenerator.Get()

}

func SpamEmailGetDBAll(app *app.App) (*[]string, error) {
	return app.SpamEmailStore.FindAll()
}
