package config

import (
	"flag"
	"os"
	"testing"
)

func TestLoadConfig_DefaultValues(t *testing.T) {
	// Удаляем переменные окружения для проверки значений по умолчанию
	os.Unsetenv("SERVICE_URL")
	os.Unsetenv("SERVICE_PORT")

	// Устанавливаем значения флагов
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	flag.Set("b", "http://localhost:8000/default-url")

	cfg := LoadConfig()

	// Проверяем значения по умолчанию
	if cfg.ServiceURL != "localhost" {
		t.Errorf("ожидалось ServiceURL = 'http://localhost', получено: %s", cfg.ServiceURL)
	}
	if cfg.ServicePort != 8080 {
		t.Errorf("ожидалось ServicePort = 8080, получено: %d", cfg.ServicePort)
	}
	if cfg.BaseShortenerURL != "http://localhost:8000/default-url" {
		t.Errorf("ожидалось BaseShortenerURL = 'http://localhost:8000/default-url', получено: %s", cfg.BaseShortenerURL)
	}
}

func TestLoadConfig_WithEnvironmentVariables(t *testing.T) {
	// Устанавливаем переменные окружения
	os.Setenv("SERVICE_URL", "http://test-service")
	os.Setenv("SERVICE_PORT", "9090")

	// Устанавливаем значения флагов
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	flag.Set("b", "http://localhost:8000/test-url")

	cfg := LoadConfig()

	// Проверяем значения из окружения
	if cfg.ServiceURL != "http://test-service" {
		t.Errorf("ожидалось ServiceURL = 'http://test-service', получено: %s", cfg.ServiceURL)
	}
	if cfg.ServicePort != 9090 {
		t.Errorf("ожидалось ServicePort = 9090, получено: %d", cfg.ServicePort)
	}
	if cfg.BaseShortenerURL != "http://localhost:8000/test-url" {
		t.Errorf("ожидалось BaseShortenerURL = 'http://localhost:8000/test-url', получено: %s", cfg.BaseShortenerURL)
	}
}

func TestServiceAddr(t *testing.T) {
	cfg := &AppConfig{
		ServiceURL:  "http://example.com",
		ServicePort: 8081,
	}

	expectedAddr := "http://example.com:8081"
	if cfg.ServiceAddr() != expectedAddr {
		t.Errorf("ожидалось ServiceAddr = '%s', получено: %s", expectedAddr, cfg.ServiceAddr())
	}
}
