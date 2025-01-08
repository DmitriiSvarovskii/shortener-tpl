package server

import (
	"net/http"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/handlers"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/storage"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	repo := storage.NewMemoryRepository()
	service := services.NewShortenerService(repo)
	handler := handlers.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Webhook)

	return &Server{
		httpServer: &http.Server{
			Handler: mux,
		},
	}
}

func (s *Server) Run(addr string) error {
	s.httpServer.Addr = addr
	return s.httpServer.ListenAndServe()
}
