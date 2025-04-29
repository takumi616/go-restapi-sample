package transform

import (
	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/model"
)

func ToModel(entity *entity.Vocabulary) *model.VocabularyInput {
	return &model.VocabularyInput{
		Title:    entity.Title,
		Meaning:  entity.Meaning,
		Sentence: entity.Sentence,
	}
}

func ToEntity(output *model.VocabularyOutput) *entity.Vocabulary {
	return &entity.Vocabulary{
		VocabularyNo: output.VocabularyNo,
		Title:        output.Title,
		Meaning:      output.Meaning,
		Sentence:     output.Sentence,
	}
}
