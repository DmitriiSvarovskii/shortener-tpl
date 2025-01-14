package server

import (
	"net/http"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/handlers"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	repo := storage.NewMemoryRepository()
	service := services.NewShortenerService(repo)
	handler := handlers.NewHandler(service)

	r := chi.NewRouter()

	r.Post("/", handler.CreateShortURLHandler)
	r.Get("/{shortURL}", handler.GetOriginalURLHandler)
	r.MethodNotAllowed(handler.MethodNotAllowedHandle)

	return &Server{
		httpServer: &http.Server{
			Handler: r,
		},
	}
}

func (s *Server) Run(addr string) error {
	s.httpServer.Addr = addr
	return s.httpServer.ListenAndServe()
}
