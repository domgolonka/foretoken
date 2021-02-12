package models

import "time"

type Spam struct {
	ID        uint
	URL       string     `db:"url"`
	Subnet    bool       `db:"subnet"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
