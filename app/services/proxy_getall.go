package services

import (
	"github.com/domgolonka/threatscraper/app"
)

func ProxyGetAll(app *app.App) []string {
	return app.ProxyGenerator.Get()

}
