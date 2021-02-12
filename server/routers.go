package server

import (
	"net/http"
	"os"

	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/lib/route"
	"github.com/domgolonka/threatscraper/ops"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router(app *app.App) http.Handler {
	r := mux.NewRouter()
	route.Attach(r, app.Config.Rooturl, PublicRoutes(app)...)

	return wrapRouter(r, app)
}

func PublicRouter(app *app.App) http.Handler {
	r := mux.NewRouter()
	route.Attach(r, app.Config.Rooturl, PublicRoutes(app)...)

	return wrapRouter(r, app)
}

func wrapRouter(r *mux.Router, app *app.App) http.Handler {
	stack := handlers.CombinedLoggingHandler(os.Stdout, r)
	app.Logger.Infof("I M HERE")
	return ops.PanicHandler(app.Reporter, stack)
}
