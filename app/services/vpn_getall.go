package services

import (
	"github.com/domgolonka/threatscraper/app"
)

func VpnGetAll(app *app.App) (*[]string, error) {
	return app.VPNGenerator.Get()

}
