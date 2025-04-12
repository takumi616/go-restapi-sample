package vocabulary

import (
	"context"
	"log/slog"
)

func (vr *VocabRepository) Delete(ctx context.Context, vocabularyNo int) (int64, error) {
	// Update the vocabulary data specified by vocabularyNo
	rowsAffected, err := vr.Persistence.Delete(ctx, vocabularyNo)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from infrastructure layer")
	}

	return rowsAffected, err
}
