package routing

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type StorageService interface {
	DeleteAllStorage(context context.Context, uid string) error
	getCollectionDocs(ctx context.Context, uid string, id string) ([]types.BasicStorageObject, error)
}

type sssRouter struct {
	service StorageService
}

func BindStorageService(route *echo.Echo, storageService StorageService) {
	self := sssRouter{service: storageService}

	route.DELETE(`/1.5/:uid`, self.deleteAllStorage)
	route.DELETE(`/1.5/:uid/storage`, self.deleteAllStorage)
	route.GET(`/1.5/:uid/storage/:collection_id:`, self.getCollection)

}

func (r *sssRouter) deleteAllStorage(c *echo.Context) error {
	uid := c.Param("uid")

	if dataError := r.service.DeleteAllStorage(c.Request().Context(), uid); dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}

	return c.String(http.StatusOK, strconv.FormatInt(time.Now().Unix(), 10))
}

func (r *sssRouter) getCollection(c *echo.Context) error {
	uid := c.Param("uid")
	collectionID := c.Param("collection_id")

	// need parse like timestamp
	//olderThen := c.QueryParamOr("older","false")
	// need parse like timestamp
	//newerThen := c.QueryParamOr("newer","false")

	isFullBso := c.QueryParamOr("full", "false")
	//sortBy := c.QueryParamOr("sort","empty")
	//limit := c.QueryParam("limit")
	//offset := c.QueryParam("offset")

	data, dataError := r.service.getCollectionDocs(c.Request().Context(), uid, collectionID)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}
	if isFullBso != "true" {
		return c.JSON(http.StatusOK, data)
	}

	out := make([]string, 0, len(data))
	for _, d := range data {
		out = append(out, d.ID)
	}
	return c.JSON(http.StatusOK, out)
}
