package vocabulary

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/request"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/transform"
)

func (vh *VocabHandler) UpdateVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the vocabularyNo from the request path
	vocabularyNo, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: "failed to convert received ID into int type"}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to the response body")
		}
		return
	}

	// Read http request body
	var req request.VocabularyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: "failed to read a request body"}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to the response body")
		}
		return
	}

	// Transform the request body into entity
	vocabulary := transform.ToEntity(&req)

	// Execute the usecase layer logic
	rowsAffected, err := vh.Usecase.UpdateVocabulary(ctx, vocabularyNo, vocabulary)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from usecase layer")
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: err.Error()}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to the response body")
		}
		return
	}

	// Write a returned result to the response body
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(response.RowsAffectedRes{RowsAffected: rowsAffected}); err != nil {
		slog.ErrorContext(ctx, "failed to write a rows affected to the response body")
		return
	}
}
