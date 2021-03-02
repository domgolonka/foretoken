package data

import (
	"fmt"
	"strings"

	"github.com/domgolonka/foretoken/app/data/postgresql"
	"github.com/domgolonka/foretoken/app/data/sqlite3"
	"github.com/domgolonka/foretoken/config"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	switch cfg.Database.Type {
	case "sqlite3":
		db, err := sqlite3.NewDB(cfg)
		if err != nil {
			return nil, err
		}
		if strings.Contains(cfg.Database.Host, ".sqlite3") {
			return db, err
		}
		// if in memory, run migrate
		err = MigrateDB(cfg, db)
		if err != nil {
			return nil, err
		}
		return db, nil
	case "postgres":
		return postgresql.NewDB(cfg)
	default:
		return nil, fmt.Errorf("unsupported database: %s", cfg.Database.Type)
	}
}

func MigrateDB(cfg *config.Config, db *sqlx.DB) (err error) {
	switch cfg.Database.Type {
	case "sqlite3":
		if db == nil {
			db, err = sqlite3.NewDB(cfg)

			if err != nil {
				return err
			}
			defer db.Close()

		}
		err = sqlite3.MigrateDB(db)
		if err != nil {
			return err
		}
		return nil
	case "postgres":
		if db == nil {
			db, err = postgresql.NewDB(cfg)
			if err != nil {
				return err
			}
			defer db.Close()
		}

		err = postgresql.MigrateDB(db)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported database: %s", cfg.Database.Type)
	}
	return nil
}
