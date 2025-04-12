package vocabulary

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vr *VocabRepository) Update(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Update the vocabulary data specified by vocabularyNo
	rowsAffected, err := vr.Persistence.Update(ctx, vocabularyNo, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return rowsAffected, err
}
