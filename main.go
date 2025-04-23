package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	vocabHandler "github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary"
	vocabRepository "github.com/takumi616/go-restapi-sample/adapter/repository/vocabulary"
	"github.com/takumi616/go-restapi-sample/infrastructure/config"
	"github.com/takumi616/go-restapi-sample/infrastructure/database"
	vocabPersistence "github.com/takumi616/go-restapi-sample/infrastructure/database/vocabulary/persistence"
	"github.com/takumi616/go-restapi-sample/infrastructure/web"
	vocabUsecase "github.com/takumi616/go-restapi-sample/usecase/vocabulary"
)

func run(ctx context.Context) error {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Get port number from config
	appInfo, err := config.GetAppInfo(ctx)
	if err != nil {
		return err
	}

	// Get DB connection information from config
	pgConnectionInfo, err := config.GetPgConnectionInfo(ctx)
	if err != nil {
		return err
	}

	// Open DB
	db, err := database.Open(ctx, pgConnectionInfo)
	if err != nil {
		return err
	}

	//Set up dependencies between layers
	vocabPersistence := vocabPersistence.New(db)
	vocabRepository := vocabRepository.New(vocabPersistence)
	vocabUsecase := vocabUsecase.New(vocabRepository)
	vocabHandler := vocabHandler.New(vocabUsecase)

	// Register handlers
	serveMux := &web.ServeMux{VocabHandler: vocabHandler}
	mux := serveMux.RegisterHandler()

	// Run server
	server := &web.Server{Port: appInfo.Port, Handler: mux}
	if err = server.Run(ctx); err != nil {
		return err
	}

	slog.InfoContext(ctx, "http server was shut down successfully")

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.ErrorContext(
			ctx, "API server could not start",
			"err", err,
		)
	}
}
