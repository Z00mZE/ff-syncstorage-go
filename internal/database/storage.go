package database

import (
	"context"
	"database/sql"
)

type Storage struct {
	conn *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{
		conn: conn,
	}
}

func (s *Storage) GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error) {
	conn, _ := s.conn.Conn(ctx)
	defer conn.Close()
	//TODO implement me
	panic("implement me")
}

func (s *Storage) GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}
func (s *Storage) GetCollectionsUsageByUID(ctx context.Context, uid uint64) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}
