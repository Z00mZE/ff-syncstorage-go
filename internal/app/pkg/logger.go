package pkg

import (
	"log/slog"

	"github.com/Z00mZE/ff-syncstorage-go/pkg/logger"
)

func NewLogger(lvl string) *slog.Logger {
	return logger.NewLogger(lvl)
}
