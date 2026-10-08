package routing

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"

	"github.com/Z00mZE/ff-syncstorage-go/internal/app/pkg/routing/storage"
	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type StorageService interface {
	DeleteAllStorage(context context.Context, uid string) error
	GetCollectionDocs(ctx context.Context, uid string, id string) ([]types.BasicStorageObject, error)
	UpsertDocuments(ctx context.Context, uid string, id string, req []types.BasicStorageObject) ([]types.BasicStorageObject, error)
}

type sssRouter struct {
	service StorageService
}

func BindStorageService(route *echo.Echo, storageService StorageService) {
	self := sssRouter{service: storageService}

	route.DELETE(`/1.5/:uid`, self.deleteAllStorage)
	route.DELETE(`/1.5/:uid/storage`, self.deleteAllStorage)
	route.GET(`/1.5/:uid/storage/:collection_id:`, self.getCollection)
	route.POST(`/1.5/:uid/storage/:collection_id:`, self.postCollection)

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

	isFullBso := c.QueryParamOr("full", "false") == "true"
	//sortBy := c.QueryParamOr("sort","empty")
	//limit := c.QueryParam("limit")
	//offset := c.QueryParam("offset")

	data, dataError := r.service.GetCollectionDocs(c.Request().Context(), uid, collectionID)
	if dataError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, dataError.Error())
	}
	if isFullBso {
		return c.JSON(http.StatusOK, data)
	}

	out := make([]string, 0, len(data))
	for _, d := range data {
		out = append(out, d.ID)
	}
	return c.JSON(http.StatusOK, out)
}

func (r *sssRouter) postCollection(c *echo.Context) error {
	uid := c.Param("uid")
	collectionID := c.Param("collection_id")
	_ = c.QueryParamOr("batch", "false")  //  Start a batch with true, or append to an existing batch using its opaque id.
	_ = c.QueryParamOr("commit", "false") //  Commit the batch identified by batch.

	var req storage.Request
	if bindError := c.Bind(&req); bindError != nil {
		return echo.NewHTTPError(http.StatusBadRequest, bindError.Error())
	}

	if len(req) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "empty collection id")
	}

	docs := make([]types.BasicStorageObject, 0, len(req))
	for _, d := range req {
		docs = append(docs, types.BasicStorageObject{
			ID:        d.Id,
			SortIndex: d.Sortindex,
			Payload:   d.Payload,
			TTL:       time.Duration(d.Ttl),
		})
	}
	result, resultError := r.service.UpsertDocuments(c.Request().Context(), uid, collectionID, docs)
	if resultError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, resultError.Error())
	}
	return c.JSON(http.StatusOK, result)
}
