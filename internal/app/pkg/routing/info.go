package routing

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type InfoService interface {
	// GetCollectionTimestamps get last-modified timestamp for every collection
	GetCollectionTimestamps(ctx context.Context, uid string) (map[string]uint64, error)
	// GetCollectionCounts get number of BSOs in each collection
	GetCollectionCounts(ctx context.Context, uid string) (map[string]uint64, error)
	// GetCollectionUsage Get data volume used by each collection in KB
	GetCollectionUsage(ctx context.Context, uid string) (map[string]uint64, error)
	// GetQuota Get current storage usage and quota in KB.
	// Two-element array [usageKB, quotaKB]; quota is null when unenforced
	GetQuota(ctx context.Context, uid string) (uint64, error)
	// GetConfiguration Get protocol and payload limits enforced by this server
	GetConfiguration(ctx context.Context, uid string) (types.Configuration, error)
}

type ssiRouter struct {
	service InfoService
}

func BindInfoService(route *echo.Echo, infoSrv InfoService) {
	self := &ssiRouter{service: infoSrv}

	route.GET(`/1.5/:uid/info/collections`, self.getCollectionTimestamps)
	route.GET(`/1.5/:uid/info/collection_counts`, self.getCollectionCounts)
	route.GET(`/1.5/:uid/info/collection_usage`, self.getCollectionUsage)
	route.GET(`/1.5/:uid/info/quota`, self.getQuota)
	route.GET(`/1.5/:uid/info/configuration`, self.getConfiguration)
}

func (r *ssiRouter) notImplemented(c *echo.Context) error {
	panic("Not Implemented")
}

func (r *ssiRouter) getCollectionTimestamps(c *echo.Context) error {
	uid := c.Param("uid")
	data, dataError := r.service.GetCollectionTimestamps(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *ssiRouter) getCollectionCounts(c *echo.Context) error {
	uid := c.Param("uid")
	data, dataError := r.service.GetCollectionCounts(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *ssiRouter) getCollectionUsage(c *echo.Context) error {
	uid := c.Param("uid")
	data, dataError := r.service.GetCollectionUsage(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *ssiRouter) getQuota(c *echo.Context) error {
	uid := c.Param("uid")
	data, dataError := r.service.GetQuota(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, []any{data, nil})
}

func (r *ssiRouter) getConfiguration(c *echo.Context) error {
	uid := c.Param("uid")
	data, dataError := r.service.GetConfiguration(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, []any{data, nil})
}
