package sqlite3

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatscraper/app/models"
	"github.com/jmoiron/sqlx"
)

type VpnStore struct {
	sqlx.Ext
}

func (db *VpnStore) FindByUrl(url string) (*models.Vpn, error) {
	vpn := models.Vpn{}
	err := sqlx.Get(db, &vpn, "SELECT * FROM vpn WHERE url = ?", url)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &vpn, nil
}

func (db *VpnStore) Find(id int) (*models.Vpn, error) {
	vpn := models.Vpn{}
	err := sqlx.Get(db, &vpn, "SELECT * FROM vpn WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &vpn, nil
}

func (db *VpnStore) FindAll() (*[]string, error) {
	vpn := []models.Vpn{}
	err := sqlx.Select(db, &vpn, "SELECT * FROM vpn")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	strings := make([]string, 0, len(vpn))
	for i := 0; i < len(vpn); i++ {
		strings = append(strings, vpn[i].URL)
	}
	return &strings, nil
}

func (db *VpnStore) Create(url string) (*models.Vpn, error) {
	now := time.Now()

	vpn := &models.Vpn{
		URL:       url,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO vpn (url,  created_at, updated_at) VALUES (:url, :created_at, :updated_at)",
		vpn,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	vpn.ID = uint(int(id))

	return vpn, nil
}

func (db *VpnStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM vpn WHERE id = ?", id)
	return ok(result, err)
}
