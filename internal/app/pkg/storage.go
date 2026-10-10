package pkg

import (
	"database/sql"

	"github.com/Z00mZE/ff-syncstorage-go/internal/database"
)

func NewStorage(conn *sql.DB) *database.Storage {
	return database.NewStorage(conn)
}
