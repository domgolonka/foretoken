package spam

type Provider interface {
	List() ([]string, []string, error)
	Name() string
}
