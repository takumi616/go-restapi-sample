package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/takumi616/go-restapi-sample/adapter/handler/request"
	"github.com/takumi616/go-restapi-sample/adapter/handler/response"
	"github.com/takumi616/go-restapi-sample/adapter/handler/transform"
	"github.com/takumi616/go-restapi-sample/adapter/handler/util"
)

const (
	CREATE_VOCABULARY_ERROR      = "Could not register a new vocabulary"
	FETCH_ALL_VOCABULARIES_ERROR = "Could not fetch the all vocabularies"
	FETCH_VOCABULARY_ERROR       = "Could not fetch the vocabulary"
	UPDATE_VOCABULARY_ERROR      = "Could not update the vocabulary"
	DELETE_VOCABULARY_ERROR      = "Could not delete the selected vocabulary"
)

type VocabHandler struct {
	Usecase VocabUsecase
}

func NewVocabHandler(vocabUsecase VocabUsecase) *VocabHandler {
	return &VocabHandler{vocabUsecase}
}

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
		slog.ErrorContext(
			ctx, "an error occurred while adding the vocabulary",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: CREATE_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, response.RowsAffectedRes{RowsAffected: rowsAffected})
}

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
	var res []*response.VocabularyRes
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
		slog.ErrorContext(
			ctx, "an error occurred while updating the vocabulary",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: UPDATE_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, response.RowsAffectedRes{RowsAffected: rowsAffected})
}

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
		slog.ErrorContext(
			ctx, "an error occurred while deleting the vocabulary",
			"err", err,
		)
		util.WriteResponse(ctx, w, http.StatusInternalServerError, response.ErrorRes{Message: DELETE_VOCABULARY_ERROR})
		return
	}

	// Write a returned result to the response body
	util.WriteResponse(ctx, w, http.StatusOK, response.RowsAffectedRes{RowsAffected: rowsAffected})
}
