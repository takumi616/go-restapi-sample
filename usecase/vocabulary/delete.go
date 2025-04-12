package vocabulary

import (
	"context"
	"log/slog"
)

func (vu *VocabUsecase) DeleteVocabulary(ctx context.Context, vocabularyNo int) (int64, error) {
	// Call repository to update the vocabulary specified by vocabularyNo
	rowsAffected, err := vu.Repository.Delete(ctx, vocabularyNo)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from adapter layer")
	}

	return rowsAffected, err
}
