package handlers

import (
	"net/http"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
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
