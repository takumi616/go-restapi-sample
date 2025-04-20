package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vu *VocabUsecase) UpdateVocabulary(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Call repository to update the vocabulary specified by vocabularyNo
	return vu.Repository.Update(ctx, vocabularyNo, vocabulary)
}
