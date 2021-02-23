package models

import "time"

type Proxy struct {
	ID        uint       `json:"-"`
	IP        string     `json:"ip" db:"ip"`
	Port      string     `json:"port" db:"port"`
	Type      string     `json:"type" db:"type"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt time.Time  `json:"-" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

func (p *Proxy) ToString() string {
	if p.Port == "" {
		return p.Type + "://" + p.IP
	}
	return p.Type + "://" + p.IP + ":" + p.Port
}
