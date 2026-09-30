package routing

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type SyncStorageService interface {
	// GetCollectionCounts возвращает список коллекций с указанием кол-ва документов BSO
	GetCollectionCounts(ctx context.Context, uid string) (map[string]uint, error)
	// GetCollectionUsage возвращает список коллекций с указанием суммарного объёма данных документов BSO (kb)
	GetCollectionUsage(ctx context.Context, uid string) (map[string]float64, error)
	// GetCollectionsTimestamps возвращает список коллекция с указанием даты последнего изменения
	GetCollectionsTimestamps(ctx context.Context, uid string) (map[string]time.Time, error)

	GetServerConfiguration(ctx context.Context, uid string) (map[string]int, error)

	GetQuotaInformation(ctx context.Context, uid string) (map[string]int, error)

	DeleteAllUserData(ctx context.Context, uid string) error
	GetCollectionBasicStorageObjectsList(
		ctx context.Context,
		uid string,
		collection string,
		bsoIds []string,
		newer time.Time,
		older time.Time,
		full bool,
		limit uint,
		offset string,
		sortOrder types.BasicStorageObjectSortOrder,
	) (map[string]int, error)

	UpsertStorageCollection(ctx context.Context, uid, collection string, data types.BasicStorageObject) error
	DeleteStorageCollection(ctx context.Context, uid, collection string, bsoIDs []string) error

	GetBasicStorageObjectID(ctx context.Context, uid, collection, bsoID string) (types.BasicStorageObject, error)
	UpsertBasicStorageObjectID(ctx context.Context, uid, collection string, data types.BasicStorageObject) error
	DeleteBasicStorageObjectID(ctx context.Context, uid, collection, bsoID string) error
}

type syncStorageInfoRouter struct {
	service SyncStorageService
}

func BindInfoService(route *echo.Echo, infoSrv SyncStorageService) {
	container := &syncStorageInfoRouter{
		service: infoSrv,
	}
	routeGroup := route.Group("/1.5/:uid/info")
	routeGroup.GET(`/collection_counts`, container.collectionCount)
	routeGroup.GET(`/collection_usage`, container.collectionUsage)
	routeGroup.GET(`/collections`, container.getCollectionsTimestamps)
	routeGroup.GET(`/configuration`, container.getServerConfiguration)

}

func (r *syncStorageInfoRouter) collectionCount(c *echo.Context) error {
	uid := c.Param("uid")
	ctx := c.Request().Context()

	data, dataError := r.service.GetCollectionCounts(ctx, uid)
	if dataError != nil {
		c.Logger().ErrorContext(ctx, dataError.Error(), slog.String("uid", uid))
		return echo.NewHTTPError(http.StatusInternalServerError, "Occurred some error getting collection counts")
	}
	return c.JSON(http.StatusOK, data)
}

func (r *syncStorageInfoRouter) collectionUsage(c *echo.Context) error {
	uid := c.Param("uid")
	ctx := c.Request().Context()
	data, dataError := r.service.GetCollectionUsage(ctx, uid)
	if dataError != nil {
		c.Logger().ErrorContext(ctx, dataError.Error(), slog.String("uid", uid))
		return echo.NewHTTPError(http.StatusInternalServerError, "Occurred some error getting collection usage")
	}
	return c.JSON(http.StatusOK, data)
}

func (r *syncStorageInfoRouter) getCollectionsTimestamps(c *echo.Context) error {
	uid := c.Param("uid")
	ctx := c.Request().Context()
	data, dataError := r.service.GetCollectionsTimestamps(ctx, uid)
	if dataError != nil {
		c.Logger().ErrorContext(ctx, dataError.Error(), slog.String("uid", uid))
		return echo.NewHTTPError(http.StatusInternalServerError, "Occurred some error getting collection usage")
	}
	out := make(map[string]int64, len(data))
	for k, v := range data {
		out[k] = v.UTC().Unix()
	}

	return c.JSON(http.StatusOK, data)
}

func (r *syncStorageInfoRouter) getServerConfiguration(c *echo.Context) error {
	return nil
}
