package web

import (
	"net/http"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary"
)

type ServeMux struct {
	VocabHandler *vocabulary.VocabHandler
}

func (s *ServeMux) RegisterHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/vocabularies", s.VocabHandler.AddVocabulary)
	mux.HandleFunc("GET /api/vocabularies", s.VocabHandler.GetVocabularyList)
	mux.HandleFunc("GET /api/vocabularies/{id}", s.VocabHandler.GetVocabularyByNo)

	return mux
}
