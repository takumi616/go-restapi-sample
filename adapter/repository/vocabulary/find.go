package vocabulary

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vr *VocabRepository) FindAll(ctx context.Context) ([]*entity.Vocabulary, error) {
	// Select all vocabulary records
	vocabularyList, err := vr.Persistence.FindAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return vocabularyList, err
}

func (vr *VocabRepository) FindByVocabNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	// Select the vocabulary specified by vocabularyNo
	vocabulary, err := vr.Persistence.FindByVocabNo(ctx, vocabularyNo)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return vocabulary, err
}
