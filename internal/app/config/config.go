package config

import (
	"flag"
)

type AppConfig struct {
	ServiceURL       string
	BaseShortenerURL string
}

func LoadConfig() *AppConfig {
	address := flag.String("a", "localhost:8080", "Адрес запуска HTTP-сервера")
	baseURL := flag.String("b", "http://localhost:8888", "Базовый адрес сокращённых URL")

	// Парсим флаги
	flag.Parse()

	// Проверяем корректность BaseURL
	if *baseURL == "" {
		panic("BaseURL не может быть пустым")
	}

	// Создаём и возвращаем конфигурацию
	return &AppConfig{
		ServiceURL:       *address,
		BaseShortenerURL: *baseURL,
	}
}
