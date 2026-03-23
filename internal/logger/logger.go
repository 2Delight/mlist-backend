package logger

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

var logger *slog.Logger

func Setup(c Config) {
	var level slog.Level
	switch c.Level {
	case LevelDebug:
		level = slog.LevelDebug
	case LevelInfo:
		level = slog.LevelInfo
	case LevelWarn:
		level = slog.LevelWarn
	case LevelError:
		level = slog.LevelError
	default:
		level = slog.LevelDebug
	}

	logger = slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: level,
			}),
			// otelslog.NewHandler("otel", otelslog.WithLoggerProvider(global.GetLoggerProvider())),
		),
	)
}
