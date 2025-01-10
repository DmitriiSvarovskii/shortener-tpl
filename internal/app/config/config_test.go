package config

import (
	"flag"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedPort    string
		expectedBaseURL string
	}{
		{
			name:            "default values",
			args:            []string{},
			expectedPort:    ":8080",
			expectedBaseURL: "http://localhost:8000/qsd54gFg",
		},
		{
			name:            "custom port",
			args:            []string{"-a", ":9090"},
			expectedPort:    ":9090",
			expectedBaseURL: "http://localhost:8000/qsd54gFg",
		},
		{
			name:            "custom base URL",
			args:            []string{"-b", "http://example.com"},
			expectedPort:    ":8080",
			expectedBaseURL: "http://example.com",
		},
		{
			name:            "custom port and base URL",
			args:            []string{"-a", ":7070", "-b", "http://test.com"},
			expectedPort:    ":7070",
			expectedBaseURL: "http://test.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сброс флагов перед каждым тестом
			flag.CommandLine = flag.NewFlagSet(tt.name, flag.ContinueOnError)
			Config = &AppConfig{} // Сброс конфигурации

			// Установка флагов
			flag.CommandLine.Parse(tt.args)

			// Вызов функции для парсинга флагов
			ParseFlags()

			// Проверка результата
			if Config.Port != tt.expectedPort {
				t.Errorf("expected Port %q, got %q", tt.expectedPort, Config.Port)
			}
			if Config.BaseShortenerURL != tt.expectedBaseURL {
				t.Errorf("expected BaseShortenerURL %q, got %q", tt.expectedBaseURL, Config.BaseShortenerURL)
			}
		})
	}
}
