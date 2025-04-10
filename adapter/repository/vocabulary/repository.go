package vocabulary

type VocabRepository struct {
	Persistence VocabPersistence
}

func New(vocabPersistence VocabPersistence) *VocabRepository {
	return &VocabRepository{vocabPersistence}
}
