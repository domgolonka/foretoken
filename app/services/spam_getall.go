package services

import (
	"github.com/domgolonka/threatscraper/app"
)

func SpamGetAll(app *app.App) []string {
	return app.SpamGenerator.Get()

}
