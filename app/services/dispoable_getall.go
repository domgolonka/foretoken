package services

import (
	"github.com/domgolonka/threatscraper/app"
)

func DisposableGetAll(app *app.App) (*[]string, error) {
	return app.DisposableGenerator.Get()

}
