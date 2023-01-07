package http

import (
	"encoding/json"
	"errors"
	"log"
	"movistar/metadata/internal/controller/metadata"
	"movistar/metadata/internal/repository"
	"net/http"
)

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	ctrl *metadata.Controller
}

// New creates a new movie metadata HTTP handler.
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMetadata handles GET /metadata requests.
func (h *Handler) GetMetadata(writer http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := request.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository has error: %v\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(writer).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
