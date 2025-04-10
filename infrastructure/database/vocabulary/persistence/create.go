package persistence

import (
	"context"
	"errors"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/vocabulary/transform"
)

func (vp *VocabPersistence) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Transform received entity into DB model
	vocabModel := transform.ToModel(vocabulary)

	// Get a DB connection from connection pool
	conn, err := vp.DB.Conn(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to establish database connection")
	}
	defer conn.Close()

	// Begin transaction
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	// Check if the vocabulary already exists
	var exists bool
	err = tx.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM vocabularies WHERE title = $1)",
		vocabModel.Title,
	).Scan(&exists)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to check existing vocabulary")
	}

	if exists {
		return 0, errors.New("vocabulary with the same title already exists")
	}

	// Execute an insert process
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO vocabularies(title, meaning, sentence) VALUES($1, $2, $3)",
		vocabModel.Title,
		vocabModel.Meaning,
		vocabModel.Sentence,
	)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to insert a new vocabulary record")
	}

	// Check rows affected number
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to get a rows affected")
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to commit transaction")
	}

	slog.InfoContext(ctx, "new vocabulary record was inserted successfully")

	return rowsAffected, nil
}
