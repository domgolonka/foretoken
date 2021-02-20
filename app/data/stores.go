package data

import (
	"fmt"

	"github.com/domgolonka/threatdefender/app/data/postgresql"

	"github.com/domgolonka/threatdefender/app/data/sqlite3"
	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

////////////////////////////
/// ipaddress
///////////////////////////
type ProxyStore interface {
	Find(id int) (*models.Proxy, error)
	FindAll() (*[]models.Proxy, error)
	FindByIP(ipaddress string) (*models.Proxy, error)
	Create(ip, port, types string) (*models.Proxy, error)
	Delete(id int) (bool, error)
}

func NewProxyStore(db sqlx.Ext) (ProxyStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.ProxyStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.ProxyStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

type VpnStore interface {
	FindByIP(ip string) (*models.Vpn, error)
	Find(id int) (*models.Vpn, error)
	FindAll() (*[]string, error)
	Create(ip string, prefix byte, iptype string, source int) (*models.Vpn, error)
	Delete(id int) (bool, error)
}

func NewVpnStore(db sqlx.Ext) (VpnStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.VpnStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.VpnStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

type SpamStore interface {
	FindByIP(ipaddress string) (*models.Spam, error)
	Find(id int) (*models.Spam, error)
	FindAll() (*[]models.Spam, error)
	FindAllIPs() (*[]string, error)
	Create(ip string, prefix byte, score int, iptype string) (*models.Spam, error)
	Delete(id int) (bool, error)
}

func NewSpamStore(db sqlx.Ext) (SpamStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.SpamStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.SpamStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

type TorStore interface {
	FindByIP(ip string) (*models.Tor, error)
	Find(id int) (*models.Tor, error)
	FindAll() (*[]string, error)
	Create(ip string, prefix byte, iptype string, source int) (*models.Tor, error)
	Delete(id int) (bool, error)
}

func NewTorStore(db sqlx.Ext) (TorStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.TorStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.TorStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

////////////////////////////
/// email
///////////////////////////
type DisposableStore interface {
	FindByEmail(email string) (*models.DisposableEmail, error)
	Find(id int) (*models.DisposableEmail, error)
	FindAll() (*[]string, error)
	Create(email string) (*models.DisposableEmail, error)
	Delete(id int) (bool, error)
}

func NewDisposableStore(db sqlx.Ext) (DisposableStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.DisposableStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.DisposableStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

type FreeEmailStore interface {
	FindByEmail(email string) (*models.FreeEmail, error)
	Find(id int) (*models.FreeEmail, error)
	FindAll() (*[]string, error)
	Create(email string) (*models.FreeEmail, error)
	Delete(id int) (bool, error)
}

func NewFreeEmailStore(db sqlx.Ext) (FreeEmailStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.FreeEmailStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.FreeEmailStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}

type SpamEmailStore interface {
	FindByEmail(email string) (*models.SpamEmail, error)
	Find(id int) (*models.SpamEmail, error)
	FindAll() (*[]string, error)
	Create(email string) (*models.SpamEmail, error)
	Delete(id int) (bool, error)
}

func NewSpamEmailStore(db sqlx.Ext) (SpamEmailStore, error) {
	switch db.DriverName() {
	case "sqlite3":
		return &sqlite3.SpamEmailStore{Ext: db}, nil
	case "postgresql":
		return &postgresql.SpamEmailStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}
}
