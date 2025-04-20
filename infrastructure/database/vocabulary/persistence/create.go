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

	// Get a DB connection from the connection pool
	conn, err := vp.DB.Conn(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to get a database connection from the connection pool",
			"err", err,
		)
		return 0, err
	}
	defer conn.Close()

	// Begin a transaction
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to begin a transaction",
			"err", err,
		)
		return 0, err
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
		slog.ErrorContext(
			ctx, "failed to check if the same vocabulary already exists",
			"err", err,
		)
		return 0, err
	}

	if exists {
		errmsg := errors.New("failed to insert the vocabulary because the same one already exists")
		slog.ErrorContext(ctx, errmsg.Error())
		return 0, errmsg
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
		slog.ErrorContext(
			ctx, "failed to insert a new vocabulary record",
			"err", err,
		)
		return 0, err
	}

	// Check rows affected number
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to get a rows affected",
			"err", err,
		)
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.ErrorContext(
			ctx, "failed to commit the transaction",
			"err", err,
		)
		return 0, err
	}

	slog.InfoContext(ctx, "new vocabulary record was inserted successfully")

	return rowsAffected, nil
}
