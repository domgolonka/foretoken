package models

import "time"

type FreeEmail struct {
	ID        uint
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
