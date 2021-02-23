package models

import "time"

type DisposableEmail struct {
	ID        uint       `json:"-"`
	Domain    string     `json:"domain" db:"domain"`
	Score     int        `json:"score" db:"score"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt time.Time  `json:"-" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}
