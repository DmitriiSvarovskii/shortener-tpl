package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *services.ShortenerService
}

func NewHandler(service *services.ShortenerService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateShortURLHandler(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "unable to read body", http.StatusInternalServerError)
		return
	}
	
	cfg := &config.AppConfig{}
	key := h.service.GenerateShortURL(string(body))
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(cfg.ServiceURL + "/" + key))
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
