package vocabulary

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vu *VocabUsecase) AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Call repository to create a new vocabulary
	rowsAffected, err := vu.Repository.Create(ctx, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from adapter layer")
	}

	return rowsAffected, err
}
