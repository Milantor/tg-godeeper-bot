package logger

import (
	"log/slog"
	"os"
)

func Init(isDebug bool) *slog.Logger {

	var handler slog.Handler
	if !isDebug {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	}
	return slog.New(handler)
}
