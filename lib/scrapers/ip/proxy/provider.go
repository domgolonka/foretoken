package proxy

import "github.com/domgolonka/foretoken/app/models"

type Provider interface {
	List() ([]models.Proxy, error)
	Name() string
	SetProxy(models.Proxy)
}
