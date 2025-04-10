package vocabulary

type VocabHandler struct {
	Usecase VocabUsecase
}

func New(vocabUsecase VocabUsecase) *VocabHandler {
	return &VocabHandler{vocabUsecase}
}
