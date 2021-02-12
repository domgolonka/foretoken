package services

import (
	"github.com/domgolonka/threatscraper/app"
)

func SpamGetCacheAll(app *app.App) []string {
	return app.SpamGenerator.Get()

}

func SpamGetDBAll(app *app.App) (*[]string, error) {
	return app.SpamStore.FindAll()
}
