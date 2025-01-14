package main

import (
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/server"
)

func main() {
	cfg := config.LoadConfig()
	srv := server.NewServer()

	if err := srv.Run(cfg.ServiceURL); err != nil {
		panic(err)
	}
}
