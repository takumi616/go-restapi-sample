package web

import (
	"net/http"

	"github.com/takumi616/go-restapi-sample/adapter/handler"
)

type ServeMux struct {
	Handler *handler.VocabHandler
}

func (s *ServeMux) RegisterHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/vocabularies", s.Handler.AddVocabulary)

	return mux
}
