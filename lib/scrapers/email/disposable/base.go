package disposable

type Provider interface {
	List() ([]string, error)
	Name() string
}
