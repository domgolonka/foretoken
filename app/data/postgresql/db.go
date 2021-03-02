package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/domgolonka/foretoken/config"
	"github.com/jmoiron/sqlx"

	// load pq library with side effects
	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	return sqlx.Connect("postgres",
		fmt.Sprintf("host=%s, port=%s, user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host, string(rune(cfg.Database.Port)), cfg.Database.Username, cfg.Database.Password, cfg.Database.Name))

}

func ok(result sql.Result, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	count, err := result.RowsAffected()
	return count > 0, err
}
