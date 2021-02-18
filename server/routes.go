package server

import (
	"github.com/domgolonka/threatdefender/app"
	"github.com/domgolonka/threatdefender/lib/route"
	"github.com/domgolonka/threatdefender/server/handlers"
)

func PublicRoutes(app *app.App) []*route.HandledRoute {
	var routes []*route.HandledRoute

	routes = append(routes,
		route.Get("/health").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetHealth(app)),
		route.Get("/ip/proxy").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetProxy(app)),
		route.Get("/ip/spam").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetSpamIP(app)),
		route.Get("/ip/vpn").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetVPN(app)),
		route.Get("/ip/tor").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetTor(app)),
		route.Get("/ip/dc-names").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetDcNames(app)),
		route.Get("/email/disposal").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetDisposal(app)),
		route.Get("/email/free").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetFree(app)),
		route.Get("/email/generic").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetGeneric(app)),
		route.Get("/email/spam").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetSpamEmail(app)),

		route.Get("/score/email").
			SecuredWith(route.Unsecured()).
			Handle(handlers.GetScoreEmail(app)),
	)

	return routes
}
