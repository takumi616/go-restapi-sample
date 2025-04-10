package vocabulary

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vr *VocabRepository) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Insert a new vocabulary data
	rowsAffected, err := vr.Persistence.Create(ctx, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return rowsAffected, err
}
