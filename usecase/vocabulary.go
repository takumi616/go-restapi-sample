package usecase

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabUsecase struct {
	Repository VocabRepository
}

func NewVocabUsecase(vocabRepository VocabRepository) *VocabUsecase {
	return &VocabUsecase{vocabRepository}
}

func (vu *VocabUsecase) AddVocabulary(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Call repository to create a new vocabulary
	return vu.Repository.Create(ctx, vocabulary)
}

func (vu *VocabUsecase) GetVocabularyList(ctx context.Context) ([]*entity.Vocabulary, error) {
	return vu.Repository.FindAll(ctx)
}

func (vu *VocabUsecase) GetVocabularyByNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	return vu.Repository.FindByVocabNo(ctx, vocabularyNo)
}

func (vu *VocabUsecase) UpdateVocabulary(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Call repository to update the vocabulary specified by vocabularyNo
	return vu.Repository.Update(ctx, vocabularyNo, vocabulary)
}

func (vu *VocabUsecase) DeleteVocabulary(ctx context.Context, vocabularyNo int) (int64, error) {
	// Call repository to update the vocabulary specified by vocabularyNo
	return vu.Repository.Delete(ctx, vocabularyNo)
}
