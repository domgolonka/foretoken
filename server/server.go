package server

import (
	"crypto/tls"

	"github.com/domgolonka/threatdefender/app"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/acme/autocert"
)

func Server(app *app.App) {

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,

		Cache: autocert.DirCache("./assets"),
	}
	// TLS Config
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// Get Certificate from Let's Encrypt
		GetCertificate: certManager.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}

	srv := fiber.New(fiber.Config{
		//Prefork:      prefork,
		ErrorHandler: Error(app),
	})
	routers(srv, app)
	if app.Config.AutoTLS {
		ln, err := tls.Listen("tcp", app.Config.PublicPort, cfg)
		if err != nil {
			panic(err)
		}
		app.Logger.Fatal(srv.Listener(ln))
	} else {
		app.Logger.Fatal(srv.Listen(app.Config.PublicPort))
	}

}
