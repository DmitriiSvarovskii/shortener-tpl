package handlers

import (
	"io"
	"net/http"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
)

type Handler struct {
	service *services.ShortenerService
}

func NewHandler(service *services.ShortenerService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Webhook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "unable to read body", http.StatusInternalServerError)
			return
		}
		key := h.service.GenerateShortURL(string(body))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://localhost:8080/" + key))
	case http.MethodGet:
		key := r.URL.Path[len("/"):]
		value, err := h.service.GetOriginalURL(key)
		if err != nil {
			http.Error(w, "key not found", http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", value)
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
