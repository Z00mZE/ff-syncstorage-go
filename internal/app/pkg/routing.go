package pkg

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/pkg/errors"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type InfoService interface {
	// GetCollectionTimestamps get last-modified timestamp for every collection
	GetCollectionTimestamps(ctx context.Context, uid uint64) (map[string]uint64, error)
	// GetCollectionCounts get number of BSOs in each collection
	GetCollectionCounts(ctx context.Context, uid uint64) (map[string]uint64, error)
	// GetCollectionUsage Get data volume used by each collection in KB
	GetCollectionUsage(ctx context.Context, uid uint64) (map[string]uint64, error)
	// GetQuota Get current orm usage and quota in KB.
	// Two-element array [usageKB, quotaKB]; quota is null when unenforced
	GetCurrentStorageUsage(ctx context.Context, uid uint64) (uint64, error)
	// GetConfiguration Get protocol and payload limits enforced by this server
	GetConfiguration(ctx context.Context, uid uint64) (types.Configuration, error)
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

func (r *ssiRouter) getCollectionTimestamps(c *echo.Context) error {
	uid, uidError := parseUint64(c.Param("uid"))
	if uidError != nil {
		return c.JSON(http.StatusBadRequest, uidError.Error())
	}

	data, dataError := r.service.GetCollectionTimestamps(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *ssiRouter) getCollectionCounts(c *echo.Context) error {
	uid, uidError := parseUint64(c.Param("uid"))
	if uidError != nil {
		return c.JSON(http.StatusBadRequest, uidError.Error())
	}

	data, dataError := r.service.GetCollectionCounts(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *ssiRouter) getCollectionUsage(c *echo.Context) error {
	uid, uidError := parseUint64(c.Param("uid"))
	if uidError != nil {
		return c.JSON(http.StatusBadRequest, uidError.Error())
	}

	data, dataError := r.service.GetCollectionUsage(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (r *ssiRouter) getQuota(c *echo.Context) error {
	uid, uidError := parseUint64(c.Param("uid"))
	if uidError != nil {
		return c.JSON(http.StatusBadRequest, uidError.Error())
	}

	data, dataError := r.service.GetCurrentStorageUsage(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, []any{data, nil})
}

func (r *ssiRouter) getConfiguration(c *echo.Context) error {
	uid, uidError := parseUint64(c.Param("uid"))
	if uidError != nil {
		return c.JSON(http.StatusBadRequest, uidError.Error())
	}

	data, dataError := r.service.GetConfiguration(c.Request().Context(), uid)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.JSON(http.StatusOK, []any{data, nil})
}

func parseUint64(in string) (uint64, error) {
	if in == "" {
		return 0, errors.New(`empty string`)
	}

	parseUid, parse := strconv.ParseUint(in, 10, 0)
	if parse != nil {
		return 0, errors.New(`invalid uint64`)
	}
	return parseUid, nil
}
