package services

import (
	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/app/models"
)

func ProxyGetCacheAll(app *app.App) []string {
	return app.ProxyGenerator.Get()

}

func ProxyGetDBAll(app *app.App) (*[]models.Proxy, error) {
	return app.ProxyStore.FindAll()
}
