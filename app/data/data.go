package data

import (
	"fmt"

	"github.com/domgolonka/threatdefender/app/data/postgresql"
	"github.com/domgolonka/threatdefender/app/data/sqlite3"
	"github.com/domgolonka/threatdefender/config"
	"github.com/jmoiron/sqlx"
	sq3 "github.com/mattn/go-sqlite3"
)

func NewDB(cfg config.Config) (*sqlx.DB, error) {
	switch cfg.DatabaseName {
	case "sqlite3":
		return sqlite3.NewDB(cfg)
	case "postgres":
		return postgresql.NewDB(cfg)
	default:
		return nil, fmt.Errorf("unsupported database: %s", cfg.DatabaseName)
	}
}

func MigrateDB(cfg config.Config) error {
	switch cfg.DatabaseName {
	case "sqlite3":
		db, err := sqlite3.NewDB(cfg)
		if err != nil {
			return err
		}
		defer db.Close()

		err = sqlite3.MigrateDB(db)
		if err != nil {
			return err
		}
		return nil
	case "postgres":
		db, err := postgresql.NewDB(cfg)
		if err != nil {
			return err
		}
		defer db.Close()

		err = postgresql.MigrateDB(db)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported database")
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
