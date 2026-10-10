package pkg

import (
	"log/slog"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/logger"
)

func NewLogger() *slog.Logger {
	return logger.NewLogger(`error`)
}
