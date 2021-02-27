package disposable

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.DisposableEmail, error)
	Name() string
}
