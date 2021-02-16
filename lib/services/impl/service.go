package impl

var (
	EmailService *emailService
)

func init() {
	EmailService = NewRegistryService()
}
