package crud

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/model"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/transform"
)

type VocabPersistence struct {
	DB *sql.DB
}

func (vp *VocabPersistence) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Transform received entity into DB model
	vocabModel := transform.ToModel(vocabulary)

	// Get a DB connection from connection pool
	conn, err := vp.DB.Conn(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return 0, errors.New("failed to establish database connection")
	}

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

func (vp *VocabPersistence) FindAll(ctx context.Context) ([]*entity.Vocabulary, error) {
	// Execute the select process
	rows, err := vp.DB.QueryContext(
		ctx,
		"SELECT vocabulary_no, title, meaning, sentence FROM vocabularies ORDER BY vocabulary_no ASC",
	)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return nil, errors.New("failed to select the expected rows")
	}
	defer rows.Close()

	// Copy the selected columns into the struct
	var vocabularyList []*entity.Vocabulary
	for rows.Next() {
		var vocabulary model.FindVocabularyOutput
		if err := rows.Scan(&vocabulary.VocabularyNo, &vocabulary.Title, &vocabulary.Meaning, &vocabulary.Sentence); err != nil {
			slog.ErrorContext(ctx, err.Error())
			return nil, errors.New("failed to copy the columns")
		}
		vocabularyList = append(vocabularyList, transform.ToEntity(&vocabulary))
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return nil, errors.New("found an error during iteration")
	}

	slog.InfoContext(ctx, "all vocabularies in database were fetched successfully")

	return vocabularyList, nil
}
