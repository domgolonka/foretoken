package spamemail

import "github.com/domgolonka/threatdefender/app/models"

type Provider interface {
	List() ([]models.SpamEmail, error)
	Name() string
}
