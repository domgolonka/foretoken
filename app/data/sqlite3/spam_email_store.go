package sqlite3

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

type SpamEmailStore struct {
	sqlx.Ext
}

func (db *SpamEmailStore) FindByURL(url string) (*models.SpamEmail, error) {
	spam := models.SpamEmail{}
	err := sqlx.Get(db, &spam, "SELECT * FROM spamemail WHERE url = ?", url)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spam, nil
}

func (db *SpamEmailStore) Find(id int) (*models.SpamEmail, error) {
	spamemail := models.SpamEmail{}
	err := sqlx.Get(db, &spamemail, "SELECT * FROM spamemail WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &spamemail, nil
}

func (db *SpamEmailStore) FindAll() (*[]string, error) {
	spam := []models.SpamEmail{}
	err := sqlx.Select(db, &spam, "SELECT * FROM spamemail")
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

func (db *SpamEmailStore) Create(url string) (*models.SpamEmail, error) {
	now := time.Now()

	spamemail := &models.SpamEmail{
		URL:       url,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT INTO spamemail (url, created_at, updated_at) VALUES (:url, :created_at, :updated_at)",
		spamemail,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	spamemail.ID = uint(int(id))

	return spamemail, nil
}

func (db *SpamEmailStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM spamemail WHERE id = ?", id)
	return ok(result, err)
}
