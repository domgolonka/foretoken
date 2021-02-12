package models

import "time"

type Proxy struct {
	ID        uint
	URL       string     `db:"url"`
	Type      string     `db:"type"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
