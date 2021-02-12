package handlers

import (
	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/app/services"
	"net/http"
)

func GetProxy(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items, err := services.ProxyGetDBAll(app)
		if err != nil {
			app.Reporter.ReportError(err)
		}
		WriteJSON(w, http.StatusOK, items)

	}
}
