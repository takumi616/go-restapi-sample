package persistence

import (
	"context"
	"log/slog"
)

func (vp *VocabPersistence) Delete(ctx context.Context, vocabularyNo int) (int64, error) {
	// Get a DB connection from connection pool
	conn, err := vp.DB.Conn(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get a database connection from the connection pool")
		return 0, err
	}
	defer conn.Close()

	// Begin a transaction
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, "failed to begin a transaction")
		return 0, err
	}
	defer tx.Rollback()

	// Execute the update process
	result, err := tx.ExecContext(
		ctx,
		"DELETE FROM vocabularies WHERE vocabulary_no = $1",
		vocabularyNo,
	)

	if err != nil {
		slog.ErrorContext(ctx, "failed to delete the vocabulary record")
		return 0, err
	}

	// Check rows affected number
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.ErrorContext(ctx, "failed to get a rows affected")
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.ErrorContext(ctx, "failed to commit the transaction")
		return 0, err
	}

	slog.InfoContext(ctx, "the vocabulary specified by vocabularyNo was deleted successfully")

	return rowsAffected, nil
}
