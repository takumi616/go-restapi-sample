package config

import (
	"errors"
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

func GetPgConnectionInfo() (*PgConnectionInfo, error) {
	pgConnectionInfo := &PgConnectionInfo{}
	if err := env.Parse(pgConnectionInfo); err != nil {
		slog.Error(err.Error())
		return nil, errors.New("failed to load postgres connection info from environment variables")
	}
	return pgConnectionInfo, nil
}
