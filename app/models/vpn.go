package models

import "time"

type Vpn struct {
	ID        uint
	IP        string     `db:"ip"`
	Port      string     `db:"port"`
	Source    string     `db:"source"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
