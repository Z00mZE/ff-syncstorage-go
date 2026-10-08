package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Z00mZE/ff-syncstorage-go/internal/repository/collections"
)

func NewCollectionsRepository(pgxPoolConn *pgxpool.Conn) *collections.Repository {
	return collections.NewRepository(pgxPoolConn)
}
