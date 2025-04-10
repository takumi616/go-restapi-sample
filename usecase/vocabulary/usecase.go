package vocabulary

type VocabUsecase struct {
	Repository VocabRepository
}

func New(vocabRepository VocabRepository) *VocabUsecase {
	return &VocabUsecase{vocabRepository}
}
