package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	httpserver "github.com/breakfront-planner/api-gateway/internal/api/http"
	"github.com/breakfront-planner/api-gateway/internal/configs"
)

type Application struct {
	conf       *configs.Configuration
	logger     *slog.Logger
	httpServer *httpserver.Server
}

func New(ctx context.Context, logger *slog.Logger) (*Application, error) {
	app := &Application{
		logger: logger,
	}
	if err := app.setConfig(); err != nil {
		return nil, fmt.Errorf("set configuration: %w", err)
	}

	if err := app.setServer(ctx, app.conf.HTTPServer); err != nil {
		return nil, fmt.Errorf("set server: %w", err)
	}
	/*

		if err := app.setGateways(); err != nil {
			app.closeGateways()
			return nil, fmt.Errorf("set gateways: %w", err)
		}
	*/

	return app, nil
}

func (a *Application) setConfig() error {
	conf, err := configs.New()
	if err != nil {
		return fmt.Errorf("new configuration: %w", err)
	}
	a.conf = conf
	return nil
}

func (a *Application) setServer(ctx context.Context, conf configs.HTTPServerConfig) error {
	srv, err := httpserver.New(ctx, conf, a.logger)
	if err != nil {
		return fmt.Errorf("create http server: %w", err)
	}
	a.httpServer = srv
	return nil
}

func (a *Application) Start(ctx context.Context) error {
	return a.httpServer.Start(ctx)
}

func (a *Application) Close(ctx context.Context) error {
	errs := []error{
		a.httpServer.Close(ctx),
	}
	return errors.Join(errs...)
}
