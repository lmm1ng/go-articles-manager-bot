package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	var opts *slog.HandlerOptions

	switch env {
	case "prod":
		opts = &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
	case "dev":
		opts = &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}
	default:
		return slog.Default()
	}

	return slog.New(slog.NewTextHandler(os.Stdout, opts))
}
