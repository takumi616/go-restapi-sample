package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"github.com/takumi616/go-restapi-sample/adapter/handler"
	"github.com/takumi616/go-restapi-sample/adapter/repository"
	"github.com/takumi616/go-restapi-sample/infrastructure/config"
	"github.com/takumi616/go-restapi-sample/infrastructure/database"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/persistence"
	"github.com/takumi616/go-restapi-sample/infrastructure/web"
	"github.com/takumi616/go-restapi-sample/usecase"
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
	vocabPersistence := persistence.NewVocabPersistence(db)
	vocabRepository := repository.NewVocabRepository(vocabPersistence)
	vocabUsecase := usecase.NewVocabUsecase(vocabRepository)
	vocabHandler := handler.NewVocabHandler(vocabUsecase)

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
