package model

type CreateVocabularyInput struct {
	Title    string
	Meaning  string
	Sentence string
}

type FindVocabularyOutput struct {
	VocabularyNo int
	Title        string
	Meaning      string
	Sentence     string
}
