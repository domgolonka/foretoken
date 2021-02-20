package vpn

import "github.com/domgolonka/threatdefender/app/models"

type Provider interface {
	List() ([]models.Vpn, error)
	Name() string
}
