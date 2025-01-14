package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type AppConfig struct {
	ServiceURL       string
	BaseShortenerURL string
}

func LoadConfig() *AppConfig {
	cfg := &AppConfig{}
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)

	flagSet.StringVar(&cfg.ServiceURL, "a", "http://localhost:8080", "base service URL")
	flagSet.StringVar(&cfg.BaseShortenerURL, "b", "http://localhost:8000", "base shortener URL")

	// Проверяем, есть ли флаги командной строки
	if len(os.Args) > 1 && !strings.Contains(os.Args[1], "-test.") {
		fmt.Println("Command-line arguments:", os.Args)
		_ = flagSet.Parse(os.Args[1:])
	}

	if envServiceURL := os.Getenv("SERVICE_URL"); envServiceURL != "" {
		cfg.ServiceURL = envServiceURL
	}
	if envBaseShortenerURL := os.Getenv("BASE_SHORTENER_URL"); envBaseShortenerURL != "" {
		cfg.BaseShortenerURL = envBaseShortenerURL
	}

	parts := strings.Split(cfg.ServiceURL, "://")
	if len(parts) > 1 {
		cfg.ServiceURL = parts[1] // оставляем только хост:порт
	}

	fmt.Println("SERVICE_URL:", cfg.ServiceURL)
	fmt.Println("BASE_SHORTENER_URL:", cfg.BaseShortenerURL)

	return cfg
}

// package config

// import (
// 	"flag"
// 	"os"
// )

// type AppConfig struct {
// 	ServiceURL       string
// 	BaseShortenerURL string
// }

// func LoadConfig() *AppConfig {
// 	cfg := &AppConfig{}

// 	flag.StringVar(&cfg.ServiceURL, "a", "localhost:8080", "base service URL")
// 	flag.StringVar(&cfg.BaseShortenerURL, "b", "http://localhost:8000", "base shortener URL")
// 	flag.Parse()

// 	if envServiceURL := os.Getenv("SERVICE_URL"); envServiceURL != "" {
// 		cfg.ServiceURL = envServiceURL
// 	}
// 	if envBaseShortenerURL := os.Getenv("BASE_SHORTENER_URL"); envBaseShortenerURL != "" {
// 		cfg.BaseShortenerURL = envBaseShortenerURL
// 	}

// 	return cfg
// }
