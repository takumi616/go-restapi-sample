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
	"github.com/takumi616/go-restapi-sample/infrastructure/database/crud"
	"github.com/takumi616/go-restapi-sample/infrastructure/web"
	"github.com/takumi616/go-restapi-sample/usecase"
)

func run(ctx context.Context) error {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Get port number from config
	appInfo, err := config.GetAppInfo()
	if err != nil {
		return err
	}

	// Get DB connection information from config
	pgConnectionInfo, err := config.GetPgConnectionInfo()
	if err != nil {
		return err
	}

	// Open DB
	db, err := database.Open(ctx, pgConnectionInfo)
	if err != nil {
		return err
	}

	// Set up dependencies between layers
	vocabPersistence := &crud.VocabPersistence{DB: db}
	vocabRepository := &repository.VocabRepository{Persistence: vocabPersistence}
	vocabUsecase := &usecase.VocabUsecase{Repository: vocabRepository}
	vocabHandler := &handler.VocabHandler{Usecase: vocabUsecase}

	// Register handlers
	serveMux := &web.ServeMux{Handler: vocabHandler}
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
		slog.ErrorContext(ctx, err.Error())
	}
}
