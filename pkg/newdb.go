package pkg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PConfig struct {
	Hostname string
	Port     string
	Username string
	Password string
	NameDB   string
	SSLmode  string
}

func InitPostgres(cfg PConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%s user =%s dbname=%s password=%s sslmode=%s",
		cfg.Hostname, cfg.Port, cfg.Username, cfg.NameDB, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ShutDown(db *sqlx.DB) error {
	return db.Close()
}
