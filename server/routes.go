package server

import (
	"github.com/domgolonka/threatscraper/app"
	"github.com/domgolonka/threatscraper/lib/route"
	"github.com/domgolonka/threatscraper/server/handlers"
)

func PublicRoutes(app *app.App) []*route.HandledRoute {
	var routes []*route.HandledRoute

	routes = append(routes,
		route.Get("/public/health").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetHealth(app)),
		route.Get("/public/ip/proxy").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetProxy(app)),

		route.Get("/public/ip/spam").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetSpam(app)),
		route.Get("/public/ip/vpn").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetVPN(app)),
		route.Get("/public/ip/dc-names").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetDcNames(app)),
		route.Get("/public/email/disposal").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetDisposal(app)),
		route.Get("/public/email/free").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetFree(app)),
	)

	return routes
}
