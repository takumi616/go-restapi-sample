package repository

import (
	"context"

	"github.com/takumi616/go-restapi-sample/entity"
)

type VocabRepository struct {
	Persistence VocabPersistence
}

func NewVocabRepository(vocabPersistence VocabPersistence) *VocabRepository {
	return &VocabRepository{vocabPersistence}
}

func (vr *VocabRepository) Create(ctx context.Context, vocabulary *entity.Vocabulary) (int64, error) {
	// Insert a new vocabulary data
	return vr.Persistence.Create(ctx, vocabulary)
}

func (vr *VocabRepository) FindAll(ctx context.Context) ([]*entity.Vocabulary, error) {
	// Select all vocabulary records
	return vr.Persistence.FindAll(ctx)
}

func (vr *VocabRepository) FindByVocabNo(ctx context.Context, vocabularyNo int) (*entity.Vocabulary, error) {
	// Select the vocabulary specified by vocabularyNo
	return vr.Persistence.FindByVocabNo(ctx, vocabularyNo)
}

func (vr *VocabRepository) Update(ctx context.Context, vocabularyNo int, vocabulary *entity.Vocabulary) (int64, error) {
	// Update the vocabulary data specified by vocabularyNo
	return vr.Persistence.Update(ctx, vocabularyNo, vocabulary)
}

func (vr *VocabRepository) Delete(ctx context.Context, vocabularyNo int) (int64, error) {
	// Delete the vocabulary data specified by vocabularyNo
	return vr.Persistence.Delete(ctx, vocabularyNo)
}
