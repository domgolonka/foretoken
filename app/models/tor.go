package models

import "time"

type Tor struct {
	ID        uint
	IP        string    `db:"ip"`
	Subnet    string    `db:"subnet"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
