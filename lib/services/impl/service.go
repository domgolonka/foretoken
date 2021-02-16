package impl

import "github.com/domgolonka/threatdefender/app"

var (
	EmailService *emailService
	IPService    *ipService
)

func InitRPCService(app *app.App) {
	EmailService = NewEmailService(app)
	IPService = NewIPService(app)
}
