package sqlite3

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

type SpamStore struct {
	sqlx.Ext
}

func (db *SpamStore) FindByIP(ipaddress string) (*models.Spam, error) {
	spam := models.Spam{}
	err := sqlx.Get(db, &spam, "SELECT * FROM spamip WHERE ip = ?", ipaddress)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spam, nil
}

func (db *SpamStore) Find(id int) (*models.Spam, error) {
	spam := models.Spam{}
	err := sqlx.Get(db, &spam, "SELECT * FROM spamip WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spam, nil
}

func (db *SpamStore) FindAll() (*[]models.Spam, error) {
	spam := []models.Spam{}
	err := sqlx.Select(db, &spam, "SELECT * FROM spamip")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spam, nil
}

func (db *SpamStore) FindAllIPs() (*[]string, error) {
	spam := []models.Spam{}
	err := sqlx.Select(db, &spam, "SELECT * FROM spamip")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	strings := make([]string, 0, len(spam))
	for i := 0; i < len(spam); i++ {
		if spam[i].Prefix > 0 {
			strings = append(strings, spam[i].IP+"/"+strconv.Itoa(int(spam[i].Prefix)))
		} else {
			strings = append(strings, spam[i].IP)
		}
	}
	return &strings, nil
}

func (db *SpamStore) Create(ip string, prefix byte, score int, iptype string) (*models.Spam, error) {
	now := time.Now()

	spam := &models.Spam{
		IP:        ip,
		Prefix:    prefix,
		Score:     score,
		Type:      iptype,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO spamip (ip, prefix, score, type, created_at, updated_at) VALUES (:ip, :prefix, :score, :type, :created_at, :updated_at)",
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
	result, err := db.Exec("DELETE FROM spamip WHERE id = ?", id)
	return ok(result, err)
}
func (db *SpamStore) DeleteOld(hour int) (bool, error) {
	result, err := db.Exec("DELETE FROM spamip WHERE created_at <= date('now','-? hour')", hour)
	return ok(result, err)
}
