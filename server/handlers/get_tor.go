package handlers

import (
	"net/http"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
)

func GetTor(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items, err := services.TorGetDBAll(app)
		if err != nil {
			app.Reporter.ReportError(err)
		}
		WriteData(w, http.StatusOK, items)
	}
}
