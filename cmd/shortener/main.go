package main

import (
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/server"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Создаём сервер с передачей конфигурации
	srv := server.NewServer(cfg)

	// Запускаем сервер
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
