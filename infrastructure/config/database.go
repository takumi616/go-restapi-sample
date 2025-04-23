package config

import (
	"context"
	"log/slog"

	"github.com/caarlos0/env"
)

type PgConnectionInfo struct {
	DBHost     string `env:"POSTGRES_HOST"`
	DBPort     string `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`
	DBSslmode  string `env:"POSTGRES_SSLMODE"`
}

func GetPgConnectionInfo(ctx context.Context) (*PgConnectionInfo, error) {
	pgConnectionInfo := &PgConnectionInfo{}
	if err := env.Parse(pgConnectionInfo); err != nil {
		slog.ErrorContext(ctx, "failed to load the postgres connection info from the environment variables")
		return pgConnectionInfo, err
	}
	return pgConnectionInfo, nil
}
