package config

import (
	"os"
	"testing"
)

func TestLoadConfig_DefaultValues(t *testing.T) {
	cfg := LoadConfig()

	if cfg.ServiceURL != "localhost:8080" {
		t.Errorf("Expected ServiceURL 'localhost:8080', got '%s'", cfg.ServiceURL)
	}
	if cfg.BaseShortenerURL != "http://localhost:8000" {
		t.Errorf("Expected BaseShortenerURL 'http://localhost:8000', got '%s'", cfg.BaseShortenerURL)
	}
}

func TestLoadConfig_EnvironmentVariables(t *testing.T) {
	os.Setenv("SERVICE_URL", "env-service-url:9000")
	os.Setenv("BASE_SHORTENER_URL", "http://env-base-shortener-url:9000")
	defer os.Unsetenv("SERVICE_URL")
	defer os.Unsetenv("BASE_SHORTENER_URL")

	cfg := LoadConfig()

	if cfg.ServiceURL != "env-service-url:9000" {
		t.Errorf("Expected ServiceURL 'env-service-url:9000', got '%s'", cfg.ServiceURL)
	}
	if cfg.BaseShortenerURL != "http://env-base-shortener-url:9000" {
		t.Errorf("Expected BaseShortenerURL 'http://env-base-shortener-url:9000', got '%s'", cfg.BaseShortenerURL)
	}
}
