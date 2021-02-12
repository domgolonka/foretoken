package dc

type Provider interface {
	List() ([]string, error)
	Name() string
}
