package response

type AddVocabularyRes struct {
	RowsAffected int64 `json:"rows_affected"`
}

type GetVocabularyRes struct {
	VocabularyNo int    `json:"vocabulary_no"`
	Title        string `json:"title"`
	Meaning      string `json:"meaning"`
	Sentence     string `json:"sentence"`
}
