package postgresql

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

type SpamStore struct {
	sqlx.Ext
}

func (db *SpamStore) FindByIP(ipaddress string) (*models.Spam, error) {
	spam := models.Spam{}
	err := sqlx.Get(db, &spam, "SELECT * FROM spam WHERE url = ?", ipaddress)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spam, nil
}

func (db *SpamStore) Find(id int) (*models.Spam, error) {
	spam := models.Spam{}
	err := sqlx.Get(db, &spam, "SELECT * FROM spam WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spam, nil
}

func (db *SpamStore) FindAll() (*[]string, error) {
	spam := []models.Spam{}
	err := sqlx.Select(db, &spam, "SELECT * FROM spam")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	strings := make([]string, 0, len(spam))
	for i := 0; i < len(spam); i++ {
		strings = append(strings, spam[i].URL)
	}
	return &strings, nil
}

func (db *SpamStore) Create(url string, sub bool) (*models.Spam, error) {
	now := time.Now()

	spam := &models.Spam{
		URL:       url,
		Subnet:    sub,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO spam (url, subnet, created_at, updated_at) VALUES (:url, :subnet, :created_at, :updated_at)",
		spam,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	spam.ID = uint(int(id))

	return spam, nil
}

func (db *SpamStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM spam WHERE id = ?", id)
	return ok(result, err)
}
