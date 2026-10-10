package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/Z00mZE/ff-syncstorage-go/internal/database/action"
)

type Storage struct {
	getCollectionTimestamps *action.GetCollectionTimestamps
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{
		getCollectionTimestamps: action.NewGetCollectionTimestamps(conn),
	}
}

func (s *Storage) GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]time.Time, error) {
	return s.getCollectionTimestamps.Run(ctx, uid)
}

func (s *Storage) GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}
func (s *Storage) GetCollectionsUsageByUID(ctx context.Context, uid uint64) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}
