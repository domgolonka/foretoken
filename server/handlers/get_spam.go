package handlers

import (
	"net/http"

	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/app/services"
)

func GetSpam(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items, err := services.SpamGetDBAll(app)
		if err != nil {
			app.Reporter.ReportError(err)
		}
		WriteData(w, http.StatusOK, items)
	}
}
