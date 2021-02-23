package models

import "time"

type Spam struct {
	ID        uint       `json:"-"`
	IP        string     `json:"ip" db:"ip"`
	Prefix    byte       `json:"prefix" db:"prefix"`
	Type      string     `json:"type" db:"type"`
	Score     int        `json:"score" db:"score"`
	CreatedAt time.Time  `json:"-" db:"created_at"`
	UpdatedAt time.Time  `json:"-" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

func (s Spam) ToString() string {
	if s.Prefix > 0 {
		return s.IP + "/" + string(s.Prefix)
	}
	return s.IP
}
