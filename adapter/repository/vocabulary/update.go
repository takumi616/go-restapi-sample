package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vr *VocabRepository) Update(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Update the vocabulary data specified by vocabularyNo
	return vr.Persistence.Update(ctx, vocabularyNo, vocabulary)
}
