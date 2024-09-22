package sl

import (
	"github.com/KrizzMU/coolback-alkol/pkg/logger/handlers/slogpretty"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func SetupLogger(env string, logger Config) *slog.Logger {
	var log *slog.Logger
	
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		logFile, err := os.OpenFile(logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log = slog.New(
			slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
		logDir := filepath.Dir(logger.Path)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}
	default:
		log = setupPrettySlog()
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
