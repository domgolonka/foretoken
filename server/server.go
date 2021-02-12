package server

import (
	"fmt"
	"github.com/domgolonka/threatscraper/app"

	"log"
	"net/http"
)

func Server(app *app.App) {
	if app.Config.PublicPort != 0 {
		go func() {
			fmt.Println(fmt.Sprintf("PUBLIC_URL: %s", fmt.Sprintf("%s:%d", app.Config.Rooturl, app.Config.PublicPort)))
			log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Config.Rooturl, app.Config.PublicPort), PublicRouter(app)))
		}()
	}
	fmt.Println(fmt.Sprintf("SERVER_URL: %s", fmt.Sprintf("%s:%d", app.Config.Rooturl, app.Config.ServerPort)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Config.Rooturl, app.Config.ServerPort), Router(app)))
}
