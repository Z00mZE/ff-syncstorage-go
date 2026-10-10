package info

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
	"github.com/Z00mZE/ff-syncstorage-go/pkg/util"
)

type Storage interface {
	GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error)
	GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error)
}
type CollectionsRepository interface {
	AllByID(ctx context.Context, ds ...uint64) ([]types.Collection, error)
}
type UserCollectionsRepository interface {
	AllByUserID(ctx context.Context, uid uint64) ([]types.UserCollection, error)
	AggregateStorageUsage(ctx context.Context, uid uint64) (uint64, error)
}

type BasicStorageObjectRepository interface {
	AggregateCountByCollectionIDForUserID(ctx context.Context, uid uint64, cids ...uint64) (map[uint64]uint64, error)
}
type Service struct {
	collectionsRepository     CollectionsRepository
	userCollectionsRepository UserCollectionsRepository
	objectsStorage            BasicStorageObjectRepository
	storage                   Storage
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
	out := make(map[string]uint64)

	userCollections, userCollectionsError := s.userCollectionsRepository.AllByUserID(ctx, uid)
	if userCollectionsError != nil {
		return nil, errors.Wrap(userCollectionsError, "error getting collections by user id")
	}

	userCollectionsIDs := make([]uint64, 0, len(userCollections))

	for _, userCollection := range userCollections {
		userCollectionsIDs = append(userCollectionsIDs, userCollection.CollectionID)
	}

	collectionsList, collectionsListError := s.collectionsRepository.AllByID(ctx, userCollectionsIDs...)

	if collectionsListError != nil {
		return nil, errors.Wrap(collectionsListError, "error getting collections")
	}

	collectionsListMap := util.SliceToMap(collectionsList, func(e types.Collection) (uint64, types.Collection) {
		return e.ID, e
	})
	userCollectionsMap := util.SliceToMap(userCollections, func(e types.UserCollection) (uint64, types.UserCollection) {
		return e.CollectionID, e
	})

	for _, record := range userCollectionsMap {
		out[collectionsListMap[record.CollectionID].Name] = userCollectionsMap[record.CollectionID].TotalBytes
	}
	return out, nil
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
