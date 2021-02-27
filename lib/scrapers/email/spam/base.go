package spamemail

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.SpamEmail, error)
	Name() string
}
