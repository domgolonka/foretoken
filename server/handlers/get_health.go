package handlers

import (
	"net/http"

	"github.com/domgolonka/threatscraper/app"
)

type health struct {
	HTTP bool `json:"http"`
}

func GetHealth(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		h := health{
			HTTP: true,
		}

		WriteJSON(w, http.StatusOK, h)
	}
}
