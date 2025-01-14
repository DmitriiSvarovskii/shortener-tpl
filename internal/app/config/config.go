package config

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type AppConfig struct {
	ServiceURL       string
	ServicePort      int
	BaseShortenerURL string
}

func (c *AppConfig) ServiceAddr() string {
	return fmt.Sprintf("%s:%d", c.ServiceURL, c.ServicePort)
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{}

	// Параметры из переменных окружения
	serviceURL, exists := os.LookupEnv("SERVICE_URL")
	if !exists {
		serviceURL = "localhost"
	}

	servicePort := 8080
	if port, exists := os.LookupEnv("SERVICE_PORT"); exists {
		_, err := fmt.Sscanf(port, "%d", &servicePort)
		if err != nil {
			log.Printf("Ошибка преобразования SERVICE_PORT: %v. Используется значение по умолчанию 8080", err)
		}
	}

	// Параметры из флагов
	flag.StringVar(&cfg.BaseShortenerURL, "b", "http://localhost:8000/qsd54gFg", "base shortener URL")
	flag.Parse()

	cfg.ServiceURL = serviceURL
	cfg.ServicePort = servicePort

	return cfg
}
