package sqlite3

import "github.com/jmoiron/sqlx"

func MigrateDB(db *sqlx.DB) error {
	migrations := []func(db *sqlx.DB) error{
		createVpn,
		createDisposable,
		createFreeEmail,
		//createGenericName,
		createSpam,
		createSpamEmail,
		createProxy,
		createTor,
	}
	for _, m := range migrations {
		if err := m(db); err != nil {
			return err
		}
	}
	return nil
}

func createVpn(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS vpn (
            id INTEGER PRIMARY KEY,
            url TEXT NOT NULL CONSTRAINT uniq UNIQUE,
            source TEXT,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}

func createDisposable(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS disposable (
            id INTEGER PRIMARY KEY,
            email TEXT NOT NULL CONSTRAINT uniq UNIQUE,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}

func createSpamEmail(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS spamemail (
            id INTEGER PRIMARY KEY,
            email TEXT NOT NULL CONSTRAINT uniq UNIQUE,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}

func createFreeEmail(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS freeemail (
            id INTEGER PRIMARY KEY,
            email TEXT NOT NULL CONSTRAINT uniq UNIQUE,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}

func createSpam(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS spam (
            id INTEGER PRIMARY KEY,
            ip TEXT CONSTRAINT uniq UNIQUE,
            prefix INT,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}

func createProxy(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS proxy (
            id INTEGER PRIMARY KEY,
            ip TEXT NOT NULL CONSTRAINT uniq UNIQUE,
             port TEXT  NOT NULL,
              type TEXT NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}

func createTor(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS tor (
            id INTEGER PRIMARY KEY,
            ip TEXT NOT NULL CONSTRAINT uniq UNIQUE,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	return err
}
