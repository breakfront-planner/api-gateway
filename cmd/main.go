package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/breakfront-planner/api-gateway/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	application, err := app.New(ctx, logger)
	if err != nil {
		logger.Error("error while inject dependencies", "error", err)
		os.Exit(1)
	}

	if err := application.Start(ctx); err != nil {
		logger.Error("running the program", "error", err)
		os.Exit(1)
	}

	closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := application.Close(closeCtx); err != nil {
		logger.Error("closing the program", "error", err)
	}
	logger.Info("application stopped")
}
