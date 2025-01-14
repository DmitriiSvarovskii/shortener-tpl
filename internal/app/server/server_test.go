package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/config"
)

// TestServer_Run тестирует запуск сервера и обработку маршрутов.
func TestServer_Run(t *testing.T) {
	cfg := &config.AppConfig{
		ServiceURL: "http://localhost:8888",
	}

	s := NewServer(cfg)

	t.Run("POST /", func(t *testing.T) {
		body := strings.NewReader("https://practicum.yandex.ru/")
		req := httptest.NewRequest(http.MethodPost, "/", body)
		w := httptest.NewRecorder()

		s.httpServer.Handler.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		responseBody, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(responseBody), cfg.ServiceURL) {
			t.Errorf("expected response to contain base URL, got %q", string(responseBody))
		}
	})

	t.Run("GET /{key}", func(t *testing.T) {
		body := strings.NewReader("https://practicum.yandex.ru/")
		req := httptest.NewRequest(http.MethodPost, "/", body)
		w := httptest.NewRecorder()
		s.httpServer.Handler.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		responseBody, _ := io.ReadAll(resp.Body)

		shortURL := strings.TrimPrefix(string(responseBody), cfg.ServiceURL)

		req = httptest.NewRequest(http.MethodGet, shortURL, nil)
		w = httptest.NewRecorder()
		s.httpServer.Handler.ServeHTTP(w, req)

		resp = w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusTemporaryRedirect {
			t.Errorf("expected status %d, got %d", http.StatusTemporaryRedirect, resp.StatusCode)
		}

		location := resp.Header.Get("Location")
		if location != "https://practicum.yandex.ru/" {
			t.Errorf("expected location %q, got %q", "https://practicum.yandex.ru/", location)
		}

		req = httptest.NewRequest(http.MethodGet, "/nonexistentKey", nil)
		w = httptest.NewRecorder()
		s.httpServer.Handler.ServeHTTP(w, req)

		resp = w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})
}
