package collections

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	conn *pgxpool.Conn
}

func NewRepository(connPool *pgxpool.Conn) *Repository {
	return &Repository{conn: connPool}
}
