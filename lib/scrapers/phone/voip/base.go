package voip

type Provider interface {
	List() ([]string, error)
	Name() string
}
