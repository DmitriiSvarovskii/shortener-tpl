package server

import (
	"net/http"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/handlers"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.AppConfig) *Server {
	repo := storage.NewMemoryRepository()
	service := services.NewShortenerService(repo)
	handler := handlers.NewHandler(service, cfg) // Передаём конфигурацию в обработчик

	r := chi.NewRouter()

	r.Post("/", handler.CreateShortURLHandler)
	r.Get("/{shortURL}", handler.GetOriginalURLHandler)
	r.MethodNotAllowed(handler.MethodNotAllowedHandle)

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.ServiceURL, // Используем адрес из конфигурации
			Handler: r,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
