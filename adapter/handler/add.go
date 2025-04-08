package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/takumi616/go-restapi-sample/adapter/handler/request"
	"github.com/takumi616/go-restapi-sample/adapter/handler/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/transform"
)

func (vh *VocabHandler) AddVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read http request body
	var req request.AddVocabularyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: "failed to read a request body"}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to response body")
		}
		return
	}

	// Transform the request body into entity
	vocabulary := transform.ToEntity(&req)

	// Execute usecase layer logic
	rowsAffected, err := vh.Usecase.AddVocabulary(ctx, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from usecase layer")
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: err.Error()}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to response body")
		}
		return
	}

	// Write a returned result to response body
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(response.AddVocabularyRes{RowsAffected: rowsAffected}); err != nil {
		slog.ErrorContext(ctx, "failed to write a rows affected to response body")
		return
	}
}
