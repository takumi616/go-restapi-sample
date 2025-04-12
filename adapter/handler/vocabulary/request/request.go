package request

type VocabularyReq struct {
	Title    string `json:"title"`
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}
