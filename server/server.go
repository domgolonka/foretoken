package server

import (
	"fmt"

	"github.com/domgolonka/threatdefender/app"

	"log"
	"net/http"
)

func Server(app *app.App) {
	if app.Config.PublicPort != 0 {
		go func() {
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.PublicPort), PublicRouter(app)))
		}()
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.ServerPort), Router(app)))
}
