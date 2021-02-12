package handlers

import (
	"net/http"

	"github.com/domgolonka/threatscraper/app"
)

func GetFree(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		h := health{
			HTTP: true,
		}

		WriteJSON(w, http.StatusOK, h)
	}
}
