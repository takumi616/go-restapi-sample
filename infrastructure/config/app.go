package config

import (
	"errors"
	"log/slog"

	"github.com/caarlos0/env"
)

type AppInfo struct {
	Port string `env:"APP_CONTAINER_PORT"`
}

func GetAppInfo() (*AppInfo, error) {
	appInfo := &AppInfo{}
	if err := env.Parse(appInfo); err != nil ||
		appInfo.Port == "" {
		slog.Error(err.Error())
		return appInfo, errors.New("failed to load app port info from environment variables")
	}

	return appInfo, nil
}
