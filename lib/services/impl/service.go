package impl

var (
	EmailService *emailService
	IPService    *ipService
)

func init() {
	EmailService = NewRegistryService()
}
