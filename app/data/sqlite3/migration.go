package sqlite3

import "github.com/jmoiron/sqlx"

func MigrateDB(db *sqlx.DB) error {
	migrations := []func(db *sqlx.DB) error{
		createVpn,
		createDisposable,
		createFreeEmail,
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
            ip TEXT NOT NULL UNIQUE,
            prefix TEXT,
            type TEXT,
            score INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS vpn_by_ip ON vpn (ip)
    `)
	return err
}

func createDisposable(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS disposable (
            id INTEGER PRIMARY KEY,
            domain TEXT NOT NULL UNIQUE,
            score INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS disposable_by_domain ON disposable (domain)
    `)
	return err
}

func createSpamEmail(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS spamemail (
            id INTEGER PRIMARY KEY,
            domain TEXT NOT NULL UNIQUE,
            score INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS spamemail_by_domain ON spamemail (domain)
    `)
	return err
}

func createFreeEmail(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS freeemail (
            id INTEGER PRIMARY KEY,
           domain TEXT NOT NULL UNIQUE,
            score INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS freeemail_by_domain ON freeemail (domain)
    `)
	return err
}

func createSpam(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS spamip (
            id INTEGER PRIMARY KEY,
            ip TEXT NOT NULL UNIQUE,
            prefix TEXT,
            type TEXT NOT NULL,
            score INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS spamip_by_ip ON spamip (ip)
    `)
	return err
}

func createProxy(db *sqlx.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS proxy (
            id INTEGER PRIMARY KEY,
            ip TEXT NOT NULL UNIQUE,
             port TEXT  NOT NULL,
              type TEXT NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS proxy_by_ip ON proxy (ip)
    `)
	return err
}

func createTor(db *sqlx.DB) error {
	_, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS tor (
            id INTEGER PRIMARY KEY,
            ip TEXT NOT NULL UNIQUE,
            prefix TEXT,
              type TEXT,
            score INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )
    `)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
        CREATE INDEX IF NOT EXISTS tor_by_ip ON tor (ip)
    `)
	return err
}
