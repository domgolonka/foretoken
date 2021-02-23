package services

import (
	"github.com/domgolonka/threatdefender/app"
)

func TorGetCacheAll(app *app.App) (*[]string, error) {
	return app.TorGenerator.Get()

}

func TorGetDBAll(app *app.App) (*[]string, error) {
	return app.TorStore.FindAllIPs()
}
