package handlers

import (
	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/app/services"
	"net/http"
	"strings"
)

func GetSpam(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items := services.SpamGetAll(app)

		stringByte := strings.Join(items, "\x0A") // x20 = space and x00 = null

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(stringByte))
	}
}
