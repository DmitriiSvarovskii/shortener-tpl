package main

import (
	"log"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
