package impl

import "github.com/domgolonka/foretoken/app"

var (
	EmailSrv *EmailService
	IPSrv    *IPService
)

func InitRPCService(app *app.App) {
	EmailSrv = NewEmailService(app)
	IPSrv = NewIPService(app)
}
