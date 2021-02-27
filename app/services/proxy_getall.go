package services

import (
	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/app/models"
)

func ProxyGetCacheAll(app *app.App) []string {
	return app.ProxyGenerator.Get()

}

func ProxyGetDBAll(app *app.App) (*[]models.Proxy, error) {
	return app.ProxyStore.FindAll()
}
