package logger

import (
	"log/slog"
	"os"

	"github.com/SemenShakhray/doc-cache/utils/logger/slogpretty"
)

func SetupLogger() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
