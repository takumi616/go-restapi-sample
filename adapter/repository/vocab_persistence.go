package repository

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabPersistence interface {
	Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error)
}
