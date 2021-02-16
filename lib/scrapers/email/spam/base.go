package spamemail

type Provider interface {
	List() ([]string, error)
	Name() string
}
