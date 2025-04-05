package transform

import (
	"github.com/takumi616/go-restapi-sample/entity"
	"github.com/takumi616/go-restapi-sample/infrastructure/database/model"
)

func ToModel(entity *entity.Vocabulary) *model.CreateVocabularyInput {
	return &model.CreateVocabularyInput{
		Title:    entity.Title,
		Meaning:  entity.Meaning,
		Sentence: entity.Sentence,
	}
}
