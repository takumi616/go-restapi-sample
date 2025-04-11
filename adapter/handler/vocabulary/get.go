package vocabulary

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/transform"
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

func (vh *VocabHandler) GetVocabularyByNo(w http.ResponseWriter, r *http.Request) {
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

	// Execute the usecase layer logic
	vocabulary, err := vh.Usecase.GetVocabularyByNo(ctx, vocabularyNo)
	if err != nil {
		slog.ErrorContext(ctx, "found an error returned from the usecase layer")
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode(&response.ErrorRes{Message: err.Error()}); err != nil {
			slog.ErrorContext(ctx, "failed to write an error message to the response body")
		}
		return
	}

	// Write the selected vocabulary to the response body
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(transform.ToResponse(vocabulary)); err != nil {
		slog.ErrorContext(ctx, "failed to write the selected vocabulary to the response body")
		return
	}
}
