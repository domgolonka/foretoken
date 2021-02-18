package sqlite3

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

type FreeEmailStore struct {
	sqlx.Ext
}

func (db *FreeEmailStore) FindByEmail(email string) (*models.FreeEmail, error) {
	freeEmail := models.FreeEmail{}
	err := sqlx.Get(db, &freeEmail, "SELECT * FROM freeemail WHERE email = ?", email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &freeEmail, nil
}

func (db *FreeEmailStore) Find(id int) (*models.FreeEmail, error) {
	freeEmail := models.FreeEmail{}
	err := sqlx.Get(db, &freeEmail, "SELECT * FROM freeemail WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &freeEmail, nil
}

func (db *FreeEmailStore) FindAll() (*[]string, error) {
	freeEmail := []models.FreeEmail{}
	err := sqlx.Select(db, &freeEmail, "SELECT * FROM freeemail")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	strings := make([]string, 0, len(freeEmail))
	for i := 0; i < len(freeEmail); i++ {
		strings = append(strings, freeEmail[i].Email)
	}
	return &strings, nil
}

func (db *FreeEmailStore) Create(email string) (*models.FreeEmail, error) {
	now := time.Now()

	freeEmail := &models.FreeEmail{
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO freeemail (email,  created_at, updated_at) VALUES (:email, :created_at, :updated_at)",
		freeEmail,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	freeEmail.ID = uint(int(id))

	return freeEmail, nil
}

func (db *FreeEmailStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM freeemail WHERE id = ?", id)
	return ok(result, err)
}
