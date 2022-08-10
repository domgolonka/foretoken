package redis

import (
	"context"
	"database/sql"
	"github.com/domgolonka/foretoken/app/models"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type ProxyStore struct {
	*redis.Client
}

func (rdb *ProxyStore) FindByIP(ipaddress string) (*models.Proxy, error) {
	proxy := models.Proxy{}
	err := rdb.Get(context.Background(), "key").Scan(&proxy)
	if err != nil {
		return &proxy, err
	}
	return &proxy, nil
}

func (rdb *ProxyStore) Find(id int) (*models.Proxy, error) {
	proxy := models.Proxy{}
	err := sqlx.Get(rdb, &proxy, "SELECT * FROM proxy WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (rdb *ProxyStore) FindAll() (*[]models.Proxy, error) {
	proxy := []models.Proxy{}
	err := sqlx.Select(rdb, &proxy, "SELECT * FROM proxy")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &proxy, nil
}

func (rdb *ProxyStore) Create(ip, port, types string) (*models.Proxy, error) {
	now := time.Now()

	proxy := &models.Proxy{
		IP:        ip,
		Port:      port,
		Type:      types,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(rdb,
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

func (rdb *ProxyStore) Delete(id int) (bool, error) {
	result, err := rdb.Exec("DELETE FROM proxy WHERE id = ?", id)
	return ok(result, err)
}

func (rdb *ProxyStore) DeleteOld(hour int) (bool, error) {
	result, err := rdb.Exec("DELETE FROM proxy WHERE created_at <= date('now','-? hour')", hour)
	return ok(result, err)
}
