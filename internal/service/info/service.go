package info

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
	"github.com/Z00mZE/ff-syncstorage-go/pkg/util"
)

type CollectionsRepository interface {
	AllByID(ctx context.Context, ds ...uint64) ([]types.Collection, error)
}
type UserCollectionsRepository interface {
	AllByUserID(ctx context.Context, uid uint64) ([]types.Collection, error)
}
type Service struct {
	collectionsRepository     CollectionsRepository
	userCollectionsRepository UserCollectionsRepository
}

func NewService(collectionRepository CollectionsRepository, userCollectionRepository UserCollectionsRepository) *Service {
	return &Service{
		collectionsRepository:     collectionRepository,
		userCollectionsRepository: userCollectionRepository,
	}
}

func (s *Service) GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error) {
	out := make(map[string]uint64)

	userCollections, userCollectionsError := s.userCollectionsRepository.AllByUserID(ctx, uid)
	if userCollectionsError != nil {
		return nil, errors.Wrap(userCollectionsError, "error getting collections by user id")
	}
	userCollectionsIDs := make([]uint64, 0, len(userCollections))
	for _, userCollection := range userCollections {
		userCollectionsIDs = append(userCollectionsIDs, userCollection.ID)
	}

	collectionsList, collectionsListError := s.collectionsRepository.AllByID(ctx, userCollectionsIDs...)

	if collectionsListError != nil {
		return nil, errors.Wrap(collectionsListError, "error getting collections")
	}
		collectionsListMap:=util.SliceToMap(collectionsList, func(e types.Collection) (uint64, string) {
			return e.ID, e.Name
		})
		userCollectionsMap:=util.SliceToMap(collectionsList, func(e types.) (uint64, time.Time) {
			return
		})

	for _,record:=range userCollections {
name:=
	}
	return out, nil
}

func (s *Service) GetCollectionCounts(ctx context.Context, uid string) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetCollectionUsage(ctx context.Context, uid string) (map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetQuota(ctx context.Context, uid string) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetConfiguration(ctx context.Context, uid string) (types.Configuration, error) {
	//TODO implement me
	panic("implement me")
}
