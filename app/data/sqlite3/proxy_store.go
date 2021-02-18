package sqlite3

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

type ProxyStore struct {
	sqlx.Ext
}

func (db *ProxyStore) FindByIP(ipaddress string) (*models.Proxy, error) {
	proxy := models.Proxy{}
	err := sqlx.Get(db, &proxy, "SELECT * FROM proxy WHERE ip = ?", ipaddress)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (db *ProxyStore) Find(id int) (*models.Proxy, error) {
	proxy := models.Proxy{}
	err := sqlx.Get(db, &proxy, "SELECT * FROM proxy WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (db *ProxyStore) FindAll() (*[]models.Proxy, error) {
	proxy := []models.Proxy{}
	err := sqlx.Select(db, &proxy, "SELECT * FROM proxy")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (db *ProxyStore) Create(ip, port, types string) (*models.Proxy, error) {
	now := time.Now()

	proxy := &models.Proxy{
		IP:        ip,
		Port:      port,
		Type:      types,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT INTO proxy (ip, port, type, created_at, updated_at) VALUES (:ip, :port, :type, :created_at, :updated_at)",
		proxy,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	proxy.ID = uint(int(id))

	return proxy, nil
}

func (db *ProxyStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM proxy WHERE id = ?", id)
	return ok(result, err)
}
