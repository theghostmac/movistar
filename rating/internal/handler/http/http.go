// Package http is a handler that handles both GET and PUT requests.
package http

import (
	"encoding/json"
	"errors"
	"log"
	"movistar/rating/internal/controller/rating"
	"movistar/rating/pkg/model"
	"net/http"
	"strconv"
)

// Handler defines a rating service controller.
type Handler struct {
	ctrl *rating.Controller
}

// New creates a new rating service HTTP handler.
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// Handle handles PUT and GET /rating requests.
func (h *Handler) Handle(writer http.ResponseWriter, request *http.Request) {
	recordID := model.RecordID(request.FormValue("id"))
	if recordID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	recordType := model.RecordType(request.FormValue("type"))
	if recordType == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	switch request.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(request.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(writer).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	case http.MethodPut:
		userID := model.UserID(request.FormValue("userID"))
		v, err := strconv.ParseFloat(request.FormValue("value"), 64)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutRating(request.Context(), recordID, recordType, &model.Rating{UserID: userID, Value: model.RatingValue(v)}); err != nil {
			log.Printf("Repository put err: %v\n", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	default:
		writer.WriteHeader(http.StatusBadRequest)
	}
}
