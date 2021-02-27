package postgresql

import (
	"database/sql"
	"time"

	"github.com/domgolonka/foretoken/app/models"
	"github.com/jmoiron/sqlx"
)

type TorStore struct {
	sqlx.Ext
}

func (db *TorStore) FindByIP(ipaddress string) (*models.Tor, error) {
	tor := models.Tor{}
	err := sqlx.Get(db, &tor, "SELECT * FROM tor WHERE ip = ?", ipaddress)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &tor, nil
}

func (db *TorStore) Find(id int) (*models.Tor, error) {
	tor := models.Tor{}
	err := sqlx.Get(db, &tor, "SELECT * FROM tor WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &tor, nil
}

func (db *TorStore) FindAll() (*[]models.Tor, error) {
	tor := []models.Tor{}
	err := sqlx.Select(db, &tor, "SELECT * FROM tor")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &tor, nil
}

func (db *TorStore) FindAllIPs() (*[]string, error) {
	tor := []models.Tor{}
	err := sqlx.Select(db, &tor, "SELECT * FROM tor")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	strings := make([]string, 0, len(tor))
	for i := 0; i < len(tor); i++ {
		strings = append(strings, tor[i].IP)
	}
	return &strings, nil
}

func (db *TorStore) Create(ip string, prefix byte, iptype string, score int) (*models.Tor, error) {
	now := time.Now()

	tor := &models.Tor{
		IP:        ip,
		Prefix:    prefix,
		Type:      iptype,
		Score:     score,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO tor (ip, prefix, type, score, created_at, updated_at) VALUES (:ip, :prefix, :type, :score, :created_at, :updated_at)",
		tor,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	tor.ID = uint(int(id))

	return tor, nil
}

func (db *TorStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM tor WHERE id = ?", id)
	return ok(result, err)
}

func (db *TorStore) DeleteOld(hour int) (bool, error) {
	result, err := db.Exec("DELETE from tor WHERE created_at <  (now() - INTERVAL '? hour' )", hour)
	return ok(result, err)
}
