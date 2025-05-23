package web

import (
	"net/http"

	"github.com/takumi616/go-restapi-sample/adapter/handler"
)

type ServeMux struct {
	VocabHandler *handler.VocabHandler
}

func (s *ServeMux) RegisterHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/vocabularies", s.VocabHandler.AddVocabulary)
	mux.HandleFunc("GET /api/vocabularies", s.VocabHandler.GetVocabularyList)
	mux.HandleFunc("GET /api/vocabularies/{id}", s.VocabHandler.GetVocabularyByNo)
	mux.HandleFunc("PUT /api/vocabularies/{id}", s.VocabHandler.UpdateVocabulary)
	mux.HandleFunc("DELETE /api/vocabularies/{id}", s.VocabHandler.DeleteVocabulary)

	return mux
}
