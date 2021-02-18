package server

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func routers(srv fiber.Router, app app.App) {
	srv.Get("/health", handlers.GetHealth())
	srv.Get("/ip/proxy", addBook(service))
	srv.Get("/ip/spam", addBook(service))
	srv.Get("/ip/vpn", addBook(service))
	srv.Get("/ip/tor", addBook(service))
	srv.Get("/ip/dc-names", addBook(service))
	srv.Get("/email/disposal", addBook(service))
	srv.Get("/email/free", addBook(service))
	srv.Get("/email/generic", addBook(service))
	srv.Get("/email/spam", addBook(service))
	srv.Get("/score/email", addBook(service))
	srv.Get("/score/ip", addBook(service))
}
