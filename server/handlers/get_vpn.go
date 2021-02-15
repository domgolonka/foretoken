package handlers

import (
	"net/http"
	"strings"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
)

func GetVPN(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items, err := services.VpnGetDBAll(app)
		if err != nil {
			WriteErrors(w, err)
		}
		stringByte := strings.Join(*items, "\x0A") // x20 = space and x00 = null

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(stringByte))
		if err != nil {
			WriteErrors(w, err)
		}
	}
}
