package models

import "time"

type Spam struct {
	ID        uint
	IP        string     `db:"ip"`
	Prefix    uint8      `db:"prefix"`
	Score     int        `db:"score"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (s Spam) ToString() string {
	if s.Prefix > 0 {
		return s.IP + "/" + string(s.Prefix)
	}
	return s.IP
}
