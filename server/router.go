package server

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func routers(srv fiber.Router, app *app.App) {

	srv.Get("/health", handlers.GetHealth())
	srv.Get("/list/ip/proxy", handlers.GetProxyIPs(app))
	srv.Get("/list/ip/spam", handlers.GetSpamIPs(app))
	srv.Get("/list/ip/vpn", handlers.GetVPNIPs(app))
	srv.Get("/list/ip/tor", handlers.GetTorIPs(app))
	srv.Get("/list/ip/dc-names", handlers.GetDCIPs(app))
	srv.Get("/list/email/disposal", handlers.GetDisasableEmails(app))
	srv.Get("/list/email/free", handlers.GetFreeEmails(app))
	srv.Get("/list/email/generic", handlers.GetGenericEmails(app))
	srv.Get("/list/email/spam", handlers.GetSpamEmails(app))
	srv.Get("/score/email/:email", handlers.GetScoreEmail(app))
	srv.Get("/score/ip/:ip", handlers.GetScoreIP(app))
	srv.Get("/validate/email/:email", handlers.GetValidateEmail(app))
	srv.Get("/email/:email", handlers.GetEmail(app))
}
