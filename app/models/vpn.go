package models

import "time"

type Vpn struct {
	ID        uint
	IP        string     `db:"ip"`
	Prefix    byte       `db:"prefix"`
	Type      string     `db:"type"`
	Score     int        `db:"score"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (v Vpn) ToString() string {
	if v.Prefix > 0 {
		return v.IP + "/" + string(v.Prefix)
	}
	return v.IP
}
