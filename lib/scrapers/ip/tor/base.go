package tor

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.Tor, error)
	Name() string
}
