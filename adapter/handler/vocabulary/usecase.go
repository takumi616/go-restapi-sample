package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabUsecase interface {
	AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error)

	GetVocabularyList(ctx context.Context) ([]*entity.Vocabulary, error)

	GetVocabularyByNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error)

	UpdateVocabulary(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error)
}
