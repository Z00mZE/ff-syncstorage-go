package pkg

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/database/sqlite"
)

func NewConnection() (*sql.DB, error) {
	conn, connError := sqlite.NewConnection(`./db.sqlite`)
	if connError != nil {
		return nil, connError
	}

	return conn, nil
}
