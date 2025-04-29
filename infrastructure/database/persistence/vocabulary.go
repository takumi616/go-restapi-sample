package persistence

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

func NewVocabPersistence(db *sql.DB) *VocabPersistence {
	return &VocabPersistence{db}
}

func (vp *VocabPersistence) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Transform received entity into DB model
	vocabModel := transform.ToModel(vocabulary)

	// Get a DB connection from the connection pool
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

	// Check if the vocabulary already exists
	var exists bool
	err = tx.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM vocabularies WHERE title = $1)",
		vocabModel.Title,
	).Scan(&exists)

	if err != nil {
		slog.ErrorContext(ctx, "failed to check if the same vocabulary already exists")
		return 0, err
	}

	if exists {
		slog.ErrorContext(ctx, "failed to insert the vocabulary because the same one already exists")
		return 0, errors.New("the same vocabulary already exists")
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
		slog.ErrorContext(ctx, "failed to insert a new vocabulary record")
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
		slog.ErrorContext(ctx, "failed to select the expected rows")
		return nil, err
	}
	defer rows.Close()

	// Copy the selected columns into the struct
	var vocabularyList []*entity.Vocabulary
	for rows.Next() {
		var vocabulary model.VocabularyOutput
		if err := rows.Scan(&vocabulary.VocabularyNo, &vocabulary.Title, &vocabulary.Meaning, &vocabulary.Sentence); err != nil {
			slog.ErrorContext(ctx, "failed to copy the columns")
			return nil, err
		}
		vocabularyList = append(vocabularyList, transform.ToEntity(&vocabulary))
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		slog.ErrorContext(ctx, "found an error during iteration")
		return nil, err
	}

	slog.InfoContext(ctx, "all vocabularies were fetched successfully")

	return vocabularyList, nil
}

func (vp *VocabPersistence) FindByVocabNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	var row model.VocabularyOutput
	if err := vp.DB.QueryRowContext(
		ctx,
		"SELECT vocabulary_no, title, meaning, sentence FROM vocabularies WHERE vocabulary_no = $1",
		vocabularyNo,
	).Scan(&row.VocabularyNo, &row.Title, &row.Meaning, &row.Sentence); err != nil {
		slog.ErrorContext(ctx, "failed to select the expected row")
		return nil, err
	}

	vocabulary := transform.ToEntity(&row)

	slog.InfoContext(ctx, "the vocabulary specified by vocabularyNo was fetched successfully")

	return vocabulary, nil
}

func (vp *VocabPersistence) Update(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Transform the received entity into DB model
	vocabModel := transform.ToModel(vocabulary)

	// Get a DB connection from the connection pool
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
		"UPDATE vocabularies SET title = $1, meaning = $2, sentence = $3 WHERE vocabulary_no = $4",
		vocabModel.Title,
		vocabModel.Meaning,
		vocabModel.Sentence,
		vocabularyNo,
	)

	if err != nil {
		slog.ErrorContext(ctx, "failed to update the vocabulary")
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

	slog.InfoContext(ctx, "the vocabulary specified by vocabularyNo was updated successfully")

	return rowsAffected, nil
}

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
