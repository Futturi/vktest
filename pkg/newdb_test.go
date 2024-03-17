package pkg

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/magiconair/properties/assert"
)

func TestInitPostgres(t *testing.T) {
	cfg := PConfig{
		Hostname: "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		NameDB:   "postgres",
		SSLmode:  "disable",
	}
	_, err1 := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%s user =%s dbname=%s password=%s sslmode=%s",
		cfg.Hostname, cfg.Port, cfg.Username, cfg.NameDB, cfg.Password, cfg.SSLmode))
	_, err := InitPostgres(cfg)
	assert.Equal(t, err1, err)
}

func TestShutdown(t *testing.T) {

	cfg := PConfig{
		Hostname: "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		NameDB:   "postgres",
		SSLmode:  "disable",
	}
	conn, _ := InitPostgres(cfg)
	err := ShutDown(conn)
	assert.Equal(t, err, ShutDown(conn))
}
