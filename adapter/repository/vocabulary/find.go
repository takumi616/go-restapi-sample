package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vr *VocabRepository) FindAll(ctx context.Context) ([]*entity.Vocabulary, error) {
	// Select all vocabulary records
	return vr.Persistence.FindAll(ctx)
}

func (vr *VocabRepository) FindByVocabNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	// Select the vocabulary specified by vocabularyNo
	return vr.Persistence.FindByVocabNo(ctx, vocabularyNo)
}
