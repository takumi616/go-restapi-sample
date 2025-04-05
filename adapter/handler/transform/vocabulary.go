package transform

import (
	"github.com/takumi616/go-restapi-sample/adapter/handler/request"
	"github.com/takumi616/go-restapi-sample/entity"
)

func ToEntity(req *request.AddVocabularyReq) *entity.Vocabulary {
	return &entity.Vocabulary{
		Title:    req.Title,
		Meaning:  req.Meaning,
		Sentence: req.Sentence,
	}
}
