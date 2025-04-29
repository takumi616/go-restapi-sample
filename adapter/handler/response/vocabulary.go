package response

type ErrorRes struct {
	Message string `json:"message"`
}

type RowsAffectedRes struct {
	RowsAffected int64 `json:"rows_affected"`
}

type VocabularyRes struct {
	VocabularyNo int    `json:"vocabulary_no"`
	Title        string `json:"title"`
	Meaning      string `json:"meaning"`
	Sentence     string `json:"sentence"`
}
