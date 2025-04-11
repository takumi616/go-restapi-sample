package vocabulary

import (
	"context"
	"log/slog"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vu *VocabUsecase) GetVocabularyList(ctx context.Context) ([]*entity.Vocabulary, error) {
	vocabularyList, err := vu.Repository.FindAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from adapter layer")
	}

	return vocabularyList, err
}

func (vu *VocabUsecase) GetVocabularyByNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	vocabulary, err := vu.Repository.FindByVocabNo(ctx, vocabularyNo)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from adapter layer")
	}

	return vocabulary, err
}
