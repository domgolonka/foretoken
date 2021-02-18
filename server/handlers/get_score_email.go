package handlers

import (
	"net/http"
	"strconv"

	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/app/services"
)

func GetScoreEmail(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		items, err := services.ScoreEmail(app, r.Form.Get("email"))
		if err != nil {
			app.Reporter.ReportError(err)
		}
		score := strconv.Itoa(int(items))
		WriteJSON(w, http.StatusOK, score)
	}
}
