package models

import "time"

type Proxy struct {
	ID        uint
	IP        string     `db:"ip"`
	Port      string     `db:"port"`
	Type      string     `db:"type"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (p *Proxy) ToString() string {
	if p.Port == "" {
		return p.Type + "://" + p.IP
	}
	return p.Type + "://" + p.IP + ":" + p.Port
}
