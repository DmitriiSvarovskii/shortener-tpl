package config

import (
	"flag"
)

type AppConfig struct {
	Port             string
	BaseShortenerURL string
}

var Config = &AppConfig{}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&Config.Port, "a", ":8080", "address and port to run server")
	flag.StringVar(&Config.BaseShortenerURL, "b", "http://localhost:8000/qsd54gFg", "base shortener URL")
	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
