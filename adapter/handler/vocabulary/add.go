package vocabulary

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/request"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/transform"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/util"
)

const CREATE_VOCABULARY_ERROR = "Could not register a new vocabulary"

func (vh *VocabHandler) AddVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Read http request body
	var req request.VocabularyReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.ErrorContext(
			ctx, "failed to read a request body",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusBadRequest, response.ErrorRes{Message: CREATE_VOCABULARY_ERROR})
		return
	}

	// Transform the request body into entity
	vocabulary := transform.ToEntity(&req)

	// Execute usecase layer logic
	rowsAffected, err := vh.Usecase.AddVocabulary(ctx, vocabulary)
	if err != nil {
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: CREATE_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, response.RowsAffectedRes{RowsAffected: rowsAffected})
}
