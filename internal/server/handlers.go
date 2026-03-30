package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/2Delight/mlist-backend/internal/logger"
	"github.com/2Delight/mlist-backend/internal/providers/models"
	"github.com/olegdayo/omniconv"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	urlQueryModelIDKey    = "id"
	urlQueryRepositoryKey = "repository"
	urlQueryVersionKey    = "version"
	errTemplate           = "{\"error\":\"%s\"}"
	noSuchModel           = "no such model"
)

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(r, w, http.StatusOK, []byte("pong"))
}

func (s *Server) getModelsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := s.modelsProvider.GetModels(r.Context())
	if err != nil {
		logger.Error(r.Context(), "failed to get models", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	data := omniconv.ConvertSlice(resp, func(m models.Model) Model { return Model(m) })
	bytes, err := json.Marshal(data)
	if err != nil {
		logger.Error(r.Context(), "failed to marshal models", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	writeResponse(r, w, http.StatusOK, bytes)
}

func (s *Server) createModelHandler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(r.Context(), "failed to read body", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	var reqModel Model
	err = json.Unmarshal(reqBytes, &reqModel)
	if err != nil {
		logger.Error(r.Context(), "failed to unmarshal body", "error", err)
		writeResponse(r, w, http.StatusBadRequest, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	respModel, err := s.modelsProvider.CreateModel(r.Context(), models.Model(reqModel))
	if err != nil {
		logger.Error(r.Context(), "failed to create model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	bytes, err := json.Marshal(models.Model(respModel))
	if err != nil {
		logger.Error(r.Context(), "failed to marshal model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	writeResponse(r, w, http.StatusCreated, bytes)
}

func (s *Server) updateModelHandler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(r.Context(), "failed to read body", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	var reqModel Model
	err = json.Unmarshal(reqBytes, &reqModel)
	if err != nil {
		logger.Error(r.Context(), "failed to unmarshal body", "error", err)
		writeResponse(r, w, http.StatusBadRequest, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	idStr := r.URL.Query().Get(urlQueryModelIDKey)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(r.Context(), "failed to parse id", "error", err)
		writeResponse(r, w, http.StatusBadRequest, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	respModel, err := s.modelsProvider.UpdateModel(r.Context(), id, models.Model(reqModel))
	switch err {
	case nil:
	case sql.ErrNoRows:
		writeResponse(r, w, http.StatusNotFound, []byte(fmt.Sprintf(errTemplate, noSuchModel)))
		return
	default:
		logger.Error(r.Context(), "failed to create model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	bytes, err := json.Marshal(models.Model(respModel))
	if err != nil {
		logger.Error(r.Context(), "failed to marshal model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	writeResponse(r, w, http.StatusOK, bytes)
}

func (s *Server) deleteModelHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get(urlQueryModelIDKey)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(r.Context(), "failed to parse id", "error", err)
		writeResponse(r, w, http.StatusBadRequest, []byte(fmt.Sprintf(errTemplate, err.Error())))
		return
	}

	err = s.modelsProvider.DeleteModel(r.Context(), id)
	switch err {
	case nil:
		fallthrough
	case sql.ErrNoRows:
		writeResponse(r, w, http.StatusNoContent, nil)
	default:
		logger.Error(r.Context(), "failed to create model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
	}
}

func (s *Server) lookupModelHandler(w http.ResponseWriter, r *http.Request) {
	repository := r.URL.Query().Get(urlQueryRepositoryKey)
	version := r.URL.Query().Get(urlQueryVersionKey)

	err := s.modelsProvider.LookupModel(r.Context(), repository, version)
	switch err {
	case nil:
		writeResponse(r, w, http.StatusOK, []byte("got model"))
	case sql.ErrNoRows:
		writeResponse(r, w, http.StatusNotFound, []byte(fmt.Sprintf(errTemplate, noSuchModel)))
	default:
		logger.Error(r.Context(), "failed to lookup model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(fmt.Sprintf(errTemplate, err.Error())))
	}
}

func writeResponse(r *http.Request, w http.ResponseWriter, statusCode int, message []byte) {
	ResponseCounter.With(prometheus.Labels{
		operationNameLabel: r.RequestURI,
		statusCodeLabel:    strconv.Itoa(statusCode),
	}).Inc()
	w.WriteHeader(statusCode)
	if len(message) > 0 {
		_, err := w.Write(message)
		if err != nil {
			logger.Error(r.Context(), "failed to write response", "error", err)
		}
	}
}
