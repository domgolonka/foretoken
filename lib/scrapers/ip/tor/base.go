package tor

type Provider interface {
	List() ([]string, error)
	Name() string
}
