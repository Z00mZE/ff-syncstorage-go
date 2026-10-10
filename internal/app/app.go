package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/Z00mZE/ff-syncstorage-go/internal/app/pkg"
	"github.com/Z00mZE/ff-syncstorage-go/internal/service/info"
)

func NewFirefoxSyncStorageApp() {
	ctx, ctxClose := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer ctxClose()

	logger := pkg.NewLogger()
	logger.Info("hello world")

	dbConn, dbConnError := pkg.NewConnection()
	if dbConnError != nil {
		logger.Error(dbConnError.Error())
		return
	}

	defer dbConn.Close()

	storage := pkg.NewStorage(dbConn)

	infoService := info.NewService(storage)
	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.RequestLogger(),
	)

	pkg.BindInfoService(e, infoService)

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	sc := echo.StartConfig{
		Address:         ":80",
		GracefulTimeout: 5 * time.Second,
	}

	if err := sc.Start(ctx, e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
