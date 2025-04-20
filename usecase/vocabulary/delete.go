package vocabulary

import (
	"context"
)

func (vu *VocabUsecase) DeleteVocabulary(ctx context.Context, vocabularyNo int) (int64, error) {
	// Call repository to update the vocabulary specified by vocabularyNo
	return vu.Repository.Delete(ctx, vocabularyNo)
}
