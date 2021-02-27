package free

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.FreeEmail, error)
	Name() string
}
