package logger

import (
	"log/slog"
	"os"
)

type Logger = slog.Logger

func NewLogger(lvl string) *Logger {
	return slog.New(getLevel(lvl))
}

func getLevel(lvl string) slog.Handler {
	var presets = map[string]slog.Handler{
		`debug`: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		`info`:  slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}),
		`warn`:  slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelWarn}),
		`error`: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelError}),
	}

	if out, isExists := presets[lvl]; isExists {
		return out
	}
	return presets[`error`]
}
