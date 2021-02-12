package proxy

type Provider interface {
	List() ([]string, error)
	Name() string
	SetProxy(string)
}
