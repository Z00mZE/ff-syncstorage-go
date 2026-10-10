package info

import (
	"context"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type Storage interface {
	GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error)
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
	return s.storage.GetCollectionTimestamps(ctx, uid)
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
