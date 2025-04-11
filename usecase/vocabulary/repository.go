package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabRepository interface {
	Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error)

	FindAll(ctx context.Context) ([]*entity.Vocabulary, error)

	FindByVocabNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error)
}
