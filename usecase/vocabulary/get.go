package vocabulary

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

func (vu *VocabUsecase) GetVocabularyList(ctx context.Context) ([]*entity.Vocabulary, error) {
	return vu.Repository.FindAll(ctx)
}

func (vu *VocabUsecase) GetVocabularyByNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	return vu.Repository.FindByVocabNo(ctx, vocabularyNo)
}
