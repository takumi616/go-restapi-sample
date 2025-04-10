package request

type AddVocabularyReq struct {
	Title    string `json:"title"`
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}
