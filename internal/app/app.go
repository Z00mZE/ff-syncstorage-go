package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/Z00mZE/ff-syncstorage-go/internal/app/pkg"
	"github.com/Z00mZE/ff-syncstorage-go/internal/app/pkg/conf"
	"github.com/Z00mZE/ff-syncstorage-go/internal/app/pkg/repository"
	"github.com/Z00mZE/ff-syncstorage-go/internal/app/pkg/routing"
	"github.com/Z00mZE/ff-syncstorage-go/internal/service/info"
	"github.com/Z00mZE/ff-syncstorage-go/pkg/types"
)

type InfoService interface {
	GetCollectionCounts(ctx context.Context, uid string) (map[string]int, error)
	GetCollectionUsage(ctx context.Context, uid string) (map[string]float64, error)
	GetCollectionsTimestamps(ctx context.Context, uid string) (map[string]int, error)

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

type TokenService interface {
	GetSyncToken(ctx context.Context, application, version string) (string, error)
}

func NewFirefoxSyncStorageApp() {
	ctx, ctxClose := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer ctxClose()

	appConf, appConfError := conf.NewAppConfig()
	fmt.Println(appConfError)

	logger := pkg.NewLogger(appConf.GetLogLevel())

	fmt.Println(appConf.GetDbDSN())

	collectionsRepository := repository.NewCollectionsRepository(nil)
	userCollectionsRepository := repository.NewUserCollectionsRepository(nil)

	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.RequestLogger(),
	)

	routing.BindInfoService(e, info.NewService(collectionsRepository, userCollectionsRepository))

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})
	if err := e.Start(":80"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
