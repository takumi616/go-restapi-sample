package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vr *VocabRepository) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Insert a new vocabulary data
	return vr.Persistence.Create(ctx, vocabulary)
}
