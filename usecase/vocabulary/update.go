package vocabulary

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vu *VocabUsecase) UpdateVocabulary(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Call repository to update the vocabulary specified by vocabularyNo
	rowsAffected, err := vu.Repository.Update(ctx, vocabularyNo, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from adapter layer")
	}

	return rowsAffected, err
}
