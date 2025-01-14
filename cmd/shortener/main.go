package main

import (
	"log"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/server"
)

func main() {
	cfg := config.LoadConfig()
	srv := server.NewServer()
	if err := srv.Run(cfg.ServiceAddr()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
