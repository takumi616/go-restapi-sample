package vocabulary

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/transform"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/util"
)

const FETCH_ALL_VOCABULARIES_ERROR = "Could not fetch the all vocabularies"
const FETCH_VOCABULARY_ERROR = "Could not fetch the vocabulary"

func (vh *VocabHandler) GetVocabularyList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Execute the usecase layer logic
	vocabularyList, err := vh.Usecase.GetVocabularyList(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx, "an error occurred while getting the vocabulary list",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: FETCH_ALL_VOCABULARIES_ERROR})
		return
	}

	// Transform the entity into the response struct
	var res []*response.GetVocabularyRes
	for _, vocabulary := range vocabularyList {
		res = append(res, transform.ToResponse(vocabulary))
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, res)
}

func (vh *VocabHandler) GetVocabularyByNo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the vocabularyNo from the request path
	vocabularyNo, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to convert received ID into int type",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusBadRequest, response.ErrorRes{Message: FETCH_VOCABULARY_ERROR})
		return
	}

	// Execute the usecase layer logic
	vocabulary, err := vh.Usecase.GetVocabularyByNo(ctx, vocabularyNo)
	if err != nil {
		slog.ErrorContext(
			ctx, "an error occurred while getting the vocabulary",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: FETCH_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, transform.ToResponse(vocabulary))
}
