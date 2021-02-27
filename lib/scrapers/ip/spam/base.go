package spam

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.Spam, error)
	Name() string
}
