package handlers

import (
	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/app/services"
	"net/http"
	"strings"
)

func GetVPN(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items, err := services.VpnGetAll(app)
		if err != nil {
			WriteErrors(w, err)
		}
		stringByte := strings.Join(*items, "\x0A") // x20 = space and x00 = null

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(stringByte))
	}
}
