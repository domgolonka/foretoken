package data

import (
	"fmt"

	"github.com/domgolonka/threatscraper/app/config"
	"github.com/domgolonka/threatscraper/app/data/postgresql"
	"github.com/domgolonka/threatscraper/app/data/sqlite3"
	"github.com/jmoiron/sqlx"
	sq3 "github.com/mattn/go-sqlite3"
)

func NewDB(cfg config.Config) (*sqlx.DB, error) {
	switch cfg.DatabaseName {
	case "sqlite3":
		return sqlite3.NewDB(cfg.Database.Host)
	case "postgres":
		return postgresql.NewDB(cfg)
	default:
		return nil, fmt.Errorf("Unsupported database: %s", cfg.DatabaseName)
	}
}

func MigrateDB(cfg config.Config) error {
	switch cfg.DatabaseName {
	case "sqlite3":
		db, err := sqlite3.NewDB(cfg.Database.Host)
		if err != nil {
			return err
		}
		defer db.Close()

		sqlite3.MigrateDB(db)
		return nil
	case "postgres":
		db, err := postgresql.NewDB(cfg)
		if err != nil {
			return err
		}
		defer db.Close()

		postgresql.MigrateDB(db)
	default:
		return fmt.Errorf("Unsupported database")
	}
	return nil
}

func IsUniquenessError(err error) bool {
	switch i := err.(type) {
	case sq3.Error:
		return i.ExtendedCode == sq3.ErrConstraintUnique
	default:
		return false
	}
}
