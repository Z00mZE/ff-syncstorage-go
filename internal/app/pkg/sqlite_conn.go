package pkg

import (
	"database/sql"
	"embed"

	_ "modernc.org/sqlite"

	"github.com/pressly/goose/v3"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/database/sqlite"
)

//go:embed ../../../migration/sqlite/*.sql
var embedMigrations embed.FS

func NewConnection() (*sql.DB, error) {
	conn, connError := sqlite.NewConnection(`./database.db`)
	if connError != nil {
		return nil, connError
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(conn, "migrations"); err != nil {
		panic(err)
	}
}
