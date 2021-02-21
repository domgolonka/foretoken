package models

import "time"

type SpamEmail struct {
	ID        uint
	Domain    string     `db:"domain"`
	Score     int        `db:"score"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
