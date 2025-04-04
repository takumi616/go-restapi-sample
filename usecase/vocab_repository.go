package usecase

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabRepository interface {
	Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error)
}
