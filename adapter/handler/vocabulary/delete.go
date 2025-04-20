package vocabulary

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/vocabulary/util"
)

const DELETE_VOCABULARY_ERROR = "Could not delete the selected vocabulary"

func (vh *VocabHandler) DeleteVocabulary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get the vocabularyNo from the request path
	vocabularyNo, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		slog.ErrorContext(
			ctx, "failed to convert received ID into int type",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusBadRequest, response.ErrorRes{Message: DELETE_VOCABULARY_ERROR})
		return
	}

	// Execute the usecase layer logic
	rowsAffected, err := vh.Usecase.DeleteVocabulary(ctx, vocabularyNo)
	if err != nil {
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: DELETE_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, response.RowsAffectedRes{RowsAffected: rowsAffected})
}
