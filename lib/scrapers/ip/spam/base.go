package spam

import "github.com/domgolonka/threatdefender/app/models"

type Provider interface {
	List() ([]models.Spam, []models.Spam, error)
	Name() string
}
