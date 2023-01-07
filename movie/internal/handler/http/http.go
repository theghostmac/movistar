package http

import (
	"encoding/json"
	"errors"
	"log"
	"movistar/movie/internal/controller/movie"
	"net/http"
)

// Handler defines a movie handler struct.
type Handler struct {
	ctrl *movie.Controller
}

// New creates a new movie HTTP handler.
func New(ctrl *movie.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

// GetMovieDetails handles the GET /movie requests.
func (h *Handler) GetMovieDetails(writer http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	details, err := h.ctrl.Get(request.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		writer.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get has error: %v\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(writer).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
