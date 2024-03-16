package pkg

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrat() error {
	m, err := migrate.New(
		"file://internal/migrate",
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	return m.Up()
}
