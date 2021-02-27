package vpn

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.Vpn, error)
	Name() string
}
