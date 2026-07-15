package app

import (
	"context"
	"log/slog"
	"os"
)

type Application struct {
	//conf       *config.Configuration
	logger *slog.Logger
	//httpServer *httpserver.Server
}

func New(ctx context.Context) (*Application, error) {
	app := &Application{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
	/*
		if err := app.setConfig(); err != nil {
			return nil, fmt.Errorf("set configuration: %w", err)
		}

			if err := app.setGateways(); err != nil {
				app.closeGateways()
				return nil, fmt.Errorf("set gateways: %w", err)
			}
			if err := app.setServer(ctx); err != nil {
				app.closeGateways()
				return nil, fmt.Errorf("set server: %w", err)
			}*/

	return app, nil
}
