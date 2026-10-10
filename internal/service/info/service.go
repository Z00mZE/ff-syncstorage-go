package info

import (
	"context"
	"time"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type Storage interface {
	GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]time.Time, error)
	GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error)
	GetCollectionsUsageByUID(ctx context.Context, uid uint64) (map[string]uint64, error)
}
type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

// GetCollectionTimestamps Get last-modified timestamp for every collection
func (s *Service) GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error) {
	data, dataError := s.storage.GetCollectionTimestamps(ctx, uid)
	if dataError != nil {
		return nil, dataError
	}
	out := make(map[string]uint64, len(data))
	for k, v := range data {
		out[k] = uint64(v.Unix())
	}
	return out, nil
}

// GetCollectionCounts Get number of BSOs in each collection
func (s *Service) GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error) {
	return s.storage.GetCollectionCounts(ctx, uid)
}

func (s *Service) GetCollectionUsage(ctx context.Context, uid uint64) (map[string]uint64, error) {
	return s.storage.GetCollectionsUsageByUID(ctx, uid)
}

func (s *Service) GetCurrentStorageUsage(ctx context.Context, uid uint64) (uint64, error) {
	data, dataError := s.GetCollectionUsage(ctx, uid)
	if dataError != nil {
		return 0, dataError
	}
	var out uint64
	for _, record := range data {
		out += record
	}
	return out, nil
}

func (s *Service) GetConfiguration(ctx context.Context, uid uint64) (types.Configuration, error) {
	return types.Configuration{}, nil
}
