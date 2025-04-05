package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/infrastructure/config"
)

func Open(ctx context.Context, pgConnectionInfo *config.PgConnectionInfo) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgConnectionInfo.DBHost, pgConnectionInfo.DBPort, pgConnectionInfo.DBUser,
		pgConnectionInfo.DBPassword, pgConnectionInfo.DBName, pgConnectionInfo.DBSslmode,
	)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return nil, errors.New("failed to open the database")
	}

	if err := db.PingContext(ctx); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return nil, errors.New("connection to the database is not alive")
	}

	slog.InfoContext(ctx, "database was opened successfully")

	return db, nil
}
