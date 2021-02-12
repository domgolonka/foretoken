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
		route.Get("/public/proxy").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetProxy(app)),

		route.Get("/public/spam").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetSpam(app)),
		route.Get("/public/vpn").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetVPN(app)),
		route.Get("/public/dc-names").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetDcNames(app)),
		route.Get("/public/disposal").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetDisposal(app)),
		route.Get("/public/free").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetFree(app)),
	)

	return routes
}
