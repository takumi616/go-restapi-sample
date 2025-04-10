package transform

import (
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/request"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/entity"
)

func ToEntity(req *request.AddVocabularyReq) *entity.Vocabulary {
	return &entity.Vocabulary{
		Title:    req.Title,
		Meaning:  req.Meaning,
		Sentence: req.Sentence,
	}
}

func ToResponse(entity *entity.Vocabulary) *response.GetVocabularyRes {
	return &response.GetVocabularyRes{
		VocabularyNo: entity.VocabularyNo,
		Title:        entity.Title,
		Meaning:      entity.Meaning,
		Sentence:     entity.Sentence,
	}
}
