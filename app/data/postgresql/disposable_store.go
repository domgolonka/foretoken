package postgresql

import (
	"database/sql"
	"time"

	"github.com/domgolonka/threatdefender/app/models"
	"github.com/jmoiron/sqlx"
)

type DisposableStore struct {
	sqlx.Ext
}

func (db *DisposableStore) FindByEmail(email string) (*models.DisposableEmail, error) {
	disposable := models.DisposableEmail{}
	err := sqlx.Get(db, &disposable, "SELECT * FROM disposable WHERE email = ?", email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &disposable, nil
}

func (db *DisposableStore) Find(id int) (*models.DisposableEmail, error) {
	disposable := models.DisposableEmail{}
	err := sqlx.Get(db, &disposable, "SELECT * FROM disposable WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &disposable, nil
}

func (db *DisposableStore) FindAll() (*[]string, error) {
	disposable := []models.DisposableEmail{}
	err := sqlx.Select(db, &disposable, "SELECT * FROM disposable")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	strings := make([]string, 0, len(disposable))
	for i := 0; i < len(disposable); i++ {
		strings = append(strings, disposable[i].Email)
	}
	return &strings, nil
}

func (db *DisposableStore) Create(email string) (*models.DisposableEmail, error) {
	now := time.Now()

	disposable := &models.DisposableEmail{
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := sqlx.NamedExec(db,
		"INSERT OR IGNORE INTO disposable (email,  created_at, updated_at) VALUES (:email, :created_at, :updated_at)",
		disposable,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}
	disposable.ID = uint(int(id))

	return disposable, nil
}

func (db *DisposableStore) Delete(id int) (bool, error) {
	result, err := db.Exec("DELETE FROM disposable WHERE id = ?", id)
	return ok(result, err)
}
