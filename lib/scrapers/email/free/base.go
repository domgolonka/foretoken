package free

import "github.com/domgolonka/threatdefender/app/models"

type Provider interface {
	List() ([]models.FreeEmail, error)
	Name() string
}
