package server

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func routers(srv fiber.Router, app *app.App) {
	srv.Get("/health", handlers.GetHealth())
	srv.Get("/ip/proxy", handlers.GetProxyIPs(app))
	srv.Get("/ip/spam", handlers.GetSpamIPs(app))
	srv.Get("/ip/vpn", handlers.GetVPNIPs(app))
	srv.Get("/ip/tor", handlers.GetTorIPs(app))
	srv.Get("/ip/dc-names", handlers.GetDCIPs(app))
	srv.Get("/email/disposal", handlers.GetDisasableEmails(app))
	srv.Get("/email/free", handlers.GetFreeEmails(app))
	srv.Get("/email/generic", handlers.GetGenericEmails(app))
	srv.Get("/email/spam", handlers.GetSpamEmails(app))
	srv.Get("/score/email", handlers.GetScoreEmail(app))
	srv.Get("/score/ip", handlers.GetScoreIP(app))
}
