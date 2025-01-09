package main

import (
	"log"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/server"
)

func main() {
	config.ParseFlags()
	srv := server.NewServer()
	if err := srv.Run(config.Config.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	// srv := server.NewServer()
	// if err := srv.Run("localhost:8080"); err != nil {
	// 	log.Fatalf("failed to start server: %v", err)
	// }
}
