package persistence

import (
	"context"
	"errors"
	"log/slog"
)

func (vp *VocabPersistence) Delete(ctx context.Context, vocabularyNo int) (int64, error) {
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
		"DELETE FROM vocabularies WHERE vocabulary_no = $1",
		vocabularyNo,
	)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to delete the vocabulary record")
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

	slog.InfoContext(ctx, "the vocabulary specified by vocabularyNo was deleted successfully")

	return rowsAffected, nil
}
