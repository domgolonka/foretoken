package impl

import "github.com/domgolonka/foretoken/app"

var (
	EmailService *emailService
	IPService    *ipService
)

func InitRPCService(app *app.App) {
	EmailService = NewEmailService(app)
	IPService = NewIPService(app)
}
