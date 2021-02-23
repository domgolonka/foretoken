package disposable

import "github.com/domgolonka/threatdefender/app/models"

type Provider interface {
	List() ([]models.DisposableEmail, error)
	Name() string
}
