package handlers

import (
	"net/http"

	"github.com/domgolonka/threatdefender/app"
)

func GetDcNames(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		h := health{
			HTTP: true,
		}

		WriteJSON(w, http.StatusOK, h)
	}
}
