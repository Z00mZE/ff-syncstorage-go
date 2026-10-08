package info

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
	"github.com/Z00mZE/ff-syncstorage-go/pkg/util"
)

type CollectionsRepository interface {
	AllByID(ctx context.Context, ds ...uint64) ([]types.Collection, error)
}
type UserCollectionsRepository interface {
	AllByUserID(ctx context.Context, uid uint64) ([]types.UserCollection, error)
}

type BasicStorageObjectRepository interface {
	AggregateCountByCollectionIDForUserID(ctx context.Context, uid uint64, cids ...uint64) (map[uint64]uint64, error)
}
type Service struct {
	collectionsRepository     CollectionsRepository
	userCollectionsRepository UserCollectionsRepository
	objectsStorage            BasicStorageObjectRepository
}

func NewService(collectionRepository CollectionsRepository, userCollectionRepository UserCollectionsRepository) *Service {
	return &Service{
		collectionsRepository:     collectionRepository,
		userCollectionsRepository: userCollectionRepository,
	}
}

// GetCollectionTimestamps Get last-modified timestamp for every collection
func (s *Service) GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error) {
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

	for _, record := range userCollections {
		out[collectionsListMap[record.CollectionID].Name] = uint64(userCollectionsMap[record.CollectionID].Modified.Unix())
	}
	return out, nil
}

// GetCollectionCounts Get number of BSOs in each collection
func (s *Service) GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error) {
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

	aggsData, aggsDataError := s.objectsStorage.AggregateCountByCollectionIDForUserID(ctx, uid, userCollectionsIDs...)
	if aggsDataError != nil {
		return nil, errors.Wrap(aggsDataError, "error getting collections by user id")
	}

	for _, record := range userCollectionsMap {
		out[collectionsListMap[record.CollectionID].Name] = aggsData[record.CollectionID]
	}

	return out, nil
}

func (s *Service) GetCollectionUsage(ctx context.Context, uid uint64) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetQuota(ctx context.Context, uid uint64) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetConfiguration(ctx context.Context, uid uint64) (types.Configuration, error) {
	//TODO implement me
	panic("implement me")
}
