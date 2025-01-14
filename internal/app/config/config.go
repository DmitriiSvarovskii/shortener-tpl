package config

import (
	"flag"
	"os"
)

type AppConfig struct {
	ServiceURL       string
	BaseShortenerURL string
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{}

	flag.StringVar(&cfg.ServiceURL, "a", "localhost:8080", "base service URL")
	flag.StringVar(&cfg.BaseShortenerURL, "b", "http://localhost:8000", "base shortener URL")
	flag.Parse()

	if envServiceURL := os.Getenv("SERVICE_URL"); envServiceURL != "" {
		cfg.ServiceURL = envServiceURL
	}
	if envBaseShortenerURL := os.Getenv("BASE_SHORTENER_URL"); envBaseShortenerURL != "" {
		cfg.BaseShortenerURL = envBaseShortenerURL
	}

	return cfg
}
