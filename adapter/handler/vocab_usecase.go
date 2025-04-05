package handler

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabUsecase interface {
	AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error)
}
