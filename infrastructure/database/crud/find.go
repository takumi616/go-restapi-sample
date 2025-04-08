package crud

import (
	"context"
	"errors"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/model"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/transform"
)

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
