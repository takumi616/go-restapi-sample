package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/takumi616/go-restapi-sample/adapter/handler/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/transform"
)

func (vh *VocabHandler) GetVocabularyList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Execute the usecase layer logic
	vocabularyList, err := vh.Usecase.GetVocabularyList(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from the usecase layer")
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: err.Error()}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to response body")
		}
		return
	}

	// Transform the entity into the response struct
	var res []*response.GetVocabularyRes
	for _, vocabulary := range vocabularyList {
		res = append(res, transform.ToResponse(vocabulary))
	}

	// Write the selected vocabularies to the response body
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(res); err != nil {
		slog.ErrorContext(ctx, "failed to write the selected vocabularies to the response body")
		return
	}
}
