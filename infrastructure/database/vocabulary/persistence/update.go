package persistence

import (
	"context"
	"errors"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/vocabulary/transform"
)

func (vp *VocabPersistence) Update(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Transform the received entity into DB model
	vocabModel := transform.ToModel(vocabulary)

	// Get a DB connection from connection pool
	conn, err := vp.DB.Conn(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to establish the database connection")
	}
	defer conn.Close()

	// Begin a transaction
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	// Execute the update process
	result, err := tx.ExecContext(
		ctx,
		"UPDATE vocabularies SET title = $1, meaning = $2, sentence = $3 WHERE vocabulary_no = $4",
		vocabModel.Title,
		vocabModel.Meaning,
		vocabModel.Sentence,
		vocabularyNo,
	)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to update the vocabulary record")
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

	slog.InfoContext(ctx, "the vocabulary specified by vocabularyNo was updated successfully")

	return rowsAffected, nil
}
