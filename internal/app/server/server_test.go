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
	// Создаём сервер
	s := NewServer()

	// Тестируем маршрут POST
	t.Run("POST /", func(t *testing.T) {
		// Создаём HTTP-запрос
		body := strings.NewReader("https://practicum.yandex.ru/")
		req := httptest.NewRequest(http.MethodPost, "/", body)
		w := httptest.NewRecorder()

		// Выполняем запрос
		s.httpServer.Handler.ServeHTTP(w, req)

		// Проверяем результат
		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}
		baseURL := "http://localhost" + config.Config.Port + "/"

		responseBody, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(responseBody), baseURL) {
			t.Errorf("expected response to contain base URL, got %q", string(responseBody))
		}
	})

	// Тестируем маршрут GET
	t.Run("GET /{key}", func(t *testing.T) {
		// Добавляем значение через POST
		body := strings.NewReader("https://practicum.yandex.ru/")
		req := httptest.NewRequest(http.MethodPost, "/", body)
		w := httptest.NewRecorder()
		s.httpServer.Handler.ServeHTTP(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		responseBody, _ := io.ReadAll(resp.Body)
		baseURL := "http://localhost" + config.Config.Port + "/"
		shortURL := strings.TrimPrefix(string(responseBody), baseURL)

		// Случай, когда ключ существует
		req = httptest.NewRequest(http.MethodGet, "/"+shortURL, nil)
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

		// Случай, когда ключ не найден
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
