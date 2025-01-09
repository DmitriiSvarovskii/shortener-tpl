package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *services.ShortenerService
}

func NewHandler(service *services.ShortenerService) *Handler {
	return &Handler{service: service}
}

// func (h *Handler) Webhook(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "unable to read body", http.StatusInternalServerError)
// 			return
// 		}
// 		key := h.service.GenerateShortURL(string(body))
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte("http://localhost:8080/" + key))
// 	case http.MethodGet:
// 		key := r.URL.Path[len("/"):]
// 		value, err := h.service.GetOriginalURL(key)
// 		if err != nil {
// 			http.Error(w, "key not found", http.StatusBadRequest)
// 			return
// 		}
// 		w.Header().Set("Location", value)
// 		w.WriteHeader(http.StatusTemporaryRedirect)
// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

func (h *Handler) CreateShortURLHandler(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "unable to read body", http.StatusInternalServerError)
		return
	}
	key := h.service.GenerateShortURL(string(body))
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("http://localhost:8080/" + key))
}

func (h *Handler) GetOriginalURLHandler(rw http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "shortURL")
	if key == "" {
		http.Error(rw, "key param is missed", http.StatusBadRequest)
		return
	}

	value, err := h.service.GetOriginalURL(key)
	if err != nil {
		http.Error(rw, "key not found", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Location", value)
	rw.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) MethodNotAllowedHandle(rw http.ResponseWriter, r *http.Request) {
	// Логируем метод и путь запроса
	fmt.Printf("Method not allowed: %s %s\n", r.Method, r.URL.Path)

	// Формируем пользовательское сообщение об ошибке
	responseMessage := fmt.Sprintf("The method '%s' is not allowed for path '%s'.", r.Method, r.URL.Path)
	rw.WriteHeader(http.StatusMethodNotAllowed) // Устанавливаем статус 405
	io.WriteString(rw, responseMessage)
}
