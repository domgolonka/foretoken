package server

import (
	"crypto/tls"
	"time"

	"github.com/domgolonka/foretoken/app"
	"github.com/domgolonka/foretoken/lib/metrics"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	fapp := fiber.New(fiber.Config{
		//Prefork:      prefork,
		ErrorHandler: Error(app),
	})
	if app.Config.RateLimit.Enabled {
		fapp.Use(limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "127.0.0.1"
			},
			Max:        app.Config.RateLimit.Max,
			Expiration: time.Duration(app.Config.RateLimit.Expiration) * time.Second,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.Get("x-forwarded-for")
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.SendStatus(fiber.StatusTooManyRequests)
			},
		}))
	}

	srv := metrics.InitPrometheus(fapp)
	routers(srv, app)
	if app.Config.AutoTLS {
		ln, err := tls.Listen("tcp", app.Config.PublicAddress, cfg)
		if err != nil {
			panic(err)
		}
		app.Logger.Fatal(fapp.Listener(ln))
	} else {
		app.Logger.Fatal(fapp.Listen(app.Config.PublicAddress))
	}

}
