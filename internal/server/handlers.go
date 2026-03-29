package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/2Delight/mlist-backend/internal/logger"
	"github.com/2Delight/mlist-backend/internal/providers/models"
	"github.com/olegdayo/omniconv"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	urlQueryModelIDKey    = "model_id"
	urlQueryRepositoryKey = "repository"
	urlQueryVersionKey    = "version"
)

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(r, w, http.StatusOK, []byte("pong"))
}

func (s *Server) getModelsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := s.modelsProvider.GetModels(r.Context())
	if err != nil {
		logger.Error(r.Context(), "failed to get models", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte("pong"))
		return
	}

	data := omniconv.ConvertSlice(resp, func(m models.Model) Model { return Model(m) })
	bytes, err := json.Marshal(data)
	if err != nil {
		logger.Error(r.Context(), "failed to marshal models", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	writeResponse(r, w, http.StatusOK, bytes)
}

func (s *Server) createModelHandler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(r.Context(), "failed to read body", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	var reqModel Model
	err = json.Unmarshal(reqBytes, &reqModel)
	if err != nil {
		logger.Error(r.Context(), "failed to unmarshal body", "error", err)
		writeResponse(r, w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	respModel, err := s.modelsProvider.CreateModel(r.Context(), models.Model(reqModel))
	if err != nil {
		logger.Error(r.Context(), "failed to create model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	bytes, err := json.Marshal(models.Model(respModel))
	if err != nil {
		logger.Error(r.Context(), "failed to marshal models", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	writeResponse(r, w, http.StatusOK, bytes)
}

func (s *Server) updateModelHandler(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(r.Context(), "failed to read body", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	var reqModel Model
	err = json.Unmarshal(reqBytes, &reqModel)
	if err != nil {
		logger.Error(r.Context(), "failed to unmarshal body", "error", err)
		writeResponse(r, w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	r.URL.Query()

	respModel, err := s.modelsProvider.CreateModel(r.Context(), models.Model(reqModel))
	if err != nil {
		logger.Error(r.Context(), "failed to create model", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	bytes, err := json.Marshal(models.Model(respModel))
	if err != nil {
		logger.Error(r.Context(), "failed to marshal models", "error", err)
		writeResponse(r, w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	writeResponse(r, w, http.StatusOK, bytes)
}

func (s *Server) deleteModelHandler(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) lookupModelHandler(w http.ResponseWriter, r *http.Request) {
}

func writeResponse(r *http.Request, w http.ResponseWriter, statusCode int, message []byte) {
	ResponseCounter.With(prometheus.Labels{
		operationNameLabel: r.RequestURI,
		statusCodeLabel:    strconv.Itoa(statusCode),
	}).Inc()
	w.WriteHeader(statusCode)
	_, err := w.Write(message)
	if err != nil {
		logger.Error(r.Context(), "failed to write response", "error", err)
	}
}
