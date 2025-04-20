package persistence

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/vocabulary/model"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/vocabulary/transform"
)

func (vp *VocabPersistence) FindAll(ctx context.Context) ([]*entity.Vocabulary, error) {
	// Execute the select process
	rows, err := vp.DB.QueryContext(
		ctx,
		"SELECT vocabulary_no, title, meaning, sentence FROM vocabularies ORDER BY vocabulary_no ASC",
	)
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to select the expected rows",
			"err", err,
		)
		return nil, err
	}
	defer rows.Close()

	// Copy the selected columns into the struct
	var vocabularyList []*entity.Vocabulary
	for rows.Next() {
		var vocabulary model.FindVocabularyOutput
		if err := rows.Scan(&vocabulary.VocabularyNo, &vocabulary.Title, &vocabulary.Meaning, &vocabulary.Sentence); err != nil {
			slog.ErrorContext(
				ctx, "failed to copy the columns",
				"err", err,
			)
			return nil, err
		}
		vocabularyList = append(vocabularyList, transform.ToEntity(&vocabulary))
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		slog.ErrorContext(
			ctx, "found an error during iteration",
			"err", err,
		)
		return nil, err
	}

	slog.InfoContext(ctx, "all vocabularies were fetched successfully")

	return vocabularyList, nil
}

func (vp *VocabPersistence) FindByVocabNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	var row model.FindVocabularyOutput
	if err := vp.DB.QueryRowContext(
		ctx,
		"SELECT vocabulary_no, title, meaning, sentence FROM vocabularies WHERE vocabulary_no = $1",
		vocabularyNo,
	).Scan(&row.VocabularyNo, &row.Title, &row.Meaning, &row.Sentence); err != nil {
		slog.ErrorContext(
			ctx, "failed to select the expected row",
			"err", err,
		)
		return nil, err
	}

	vocabulary := transform.ToEntity(&row)

	slog.InfoContext(ctx, "the vocabulary specified by vocabularyNo was fetched successfully")

	return vocabulary, nil
}
