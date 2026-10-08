package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Z00mZE/ff-syncstorage-go/internal/repository/user_collections"
)

func NewUserCollectionsRepository(pgxPoolConn *pgxpool.Conn) *user_collections.Repository {
	return user_collections.NewRepository(pgxPoolConn)
}
