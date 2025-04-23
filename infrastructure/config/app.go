package config

import (
	"context"
	"errors"
	"log/slog"

	"github.com/caarlos0/env"
)

type AppInfo struct {
	Port string `env:"APP_CONTAINER_PORT"`
}

func GetAppInfo(ctx context.Context) (*AppInfo, error) {
	appInfo := &AppInfo{}
	if err := env.Parse(appInfo); err != nil {
		slog.ErrorContext(ctx, "failed to load the app port info from environment the variables")
		return appInfo, err
	}

	if appInfo.Port == "" {
		slog.ErrorContext(ctx, "failed to get an expected port number")
		return appInfo, errors.New("port number must be set")
	}

	return appInfo, nil
}
