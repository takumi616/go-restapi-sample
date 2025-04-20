package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vu *VocabUsecase) AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Call repository to create a new vocabulary
	return vu.Repository.Create(ctx, vocabulary)
}
