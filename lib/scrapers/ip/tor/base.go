package tor

import "github.com/domgolonka/threatdefender/app/models"

type Provider interface {
	List() ([]models.Tor, error)
	Name() string
}
