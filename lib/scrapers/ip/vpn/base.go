package vpn

type Provider interface {
	List() ([]string, error)
	Name() string
}
