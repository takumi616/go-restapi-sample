package vocabulary

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/request"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/transform"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/util"
)

const UPDATE_VOCABULARY_ERROR = "Could not update the vocabulary"

func (vh *VocabHandler) UpdateVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the vocabularyNo from the request path
	vocabularyNo, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to convert received ID into int type",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusBadRequest, response.ErrorRes{Message: UPDATE_VOCABULARY_ERROR})
		return
	}

	// Read http request body
	var req request.VocabularyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(
			ctx, "failed to read a request body",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusBadRequest, response.ErrorRes{Message: UPDATE_VOCABULARY_ERROR})
		return
	}

	// Transform the request body into entity
	vocabulary := transform.ToEntity(&req)

	// Execute the usecase layer logic
	rowsAffected, err := vh.Usecase.UpdateVocabulary(ctx, vocabularyNo, vocabulary)
	if err != nil {
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: UPDATE_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, response.RowsAffectedRes{RowsAffected: rowsAffected})
}
