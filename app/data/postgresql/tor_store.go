package postgresql

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatscraper/app/models"
	"github.com/jmoiron/sqlx"
)

type TorStore struct {
	sqlx.Ext
}

func (db *TorStore) FindByURL(url string) (*models.Tor, error) {
	tor := models.Tor{}
	err := sqlx.Get(db, &tor, "SELECT * FROM tor WHERE url = ?", url)
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

func (db *TorStore) FindAll() (*[]string, error) {
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

func (db *TorStore) Create(ip string) (*models.Tor, error) {
	now := time.Now()

	tor := &models.Tor{
		IP:        ip,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO tor (url, subnet, created_at, updated_at) VALUES (:url, :subnet, :created_at, :updated_at)",
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
