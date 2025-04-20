package vocabulary

import (
	"context"
)

func (vr *VocabRepository) Delete(ctx context.Context, vocabularyNo int) (int64, error) {
	// Delete the vocabulary data specified by vocabularyNo
	return vr.Persistence.Delete(ctx, vocabularyNo)
}
