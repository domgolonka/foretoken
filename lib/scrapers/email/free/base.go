package free

type Provider interface {
	List() ([]string, error)
	Name() string
}
