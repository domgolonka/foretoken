package server

import (
	"crypto/tls"
	"fmt"

	"golang.org/x/crypto/acme/autocert"

	"github.com/domgolonka/threatdefender/app"

	"log"
	"net/http"
)

func Server(app *app.App) {
	// create the autocert.Manager with domains and path to the cache
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("./assets"),
	}
	server := &http.Server{
		Addr:    ":https",
		Handler: PublicRouter(app),
		TLSConfig: &tls.Config{
			MinVersion:     tls.VersionTLS12,
			GetCertificate: certManager.GetCertificate,
		},
	}
	if app.Config.PublicPort != 0 {
		go func() {
			if app.Config.AutoTLS {
				log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.PublicPort), certManager.HTTPHandler(PublicRouter(app))))
			} else {
				log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.PublicPort), PublicRouter(app)))
			}

		}()
	}
	go log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.ServerPort), Router(app)))
	if app.Config.AutoTLS {
		log.Fatal(server.ListenAndServeTLS("", ""))
	}

}
