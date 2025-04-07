package repository

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabRepository struct {
	Persistence VocabPersistence
}

func (vr *VocabRepository) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Insert a new vocabulary data
	rowsAffected, err := vr.Persistence.Create(ctx, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return rowsAffected, err
}

func (vr *VocabRepository) FindAll(ctx context.Context) ([]*entity.Vocabulary, error) {
	// Select all vocabulary records
	vocabularyList, err := vr.Persistence.FindAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return vocabularyList, err
}
