package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/services"
	"github.com/go-chi/chi/v5"
)

// MockRepository имитирует поведение хранилища.
type MockRepository struct {
	data map[string]string
}

func NewMockRepository() *MockRepository {
	return &MockRepository{data: make(map[string]string)}
}

func (m *MockRepository) Save(key, value string) {
	m.data[key] = value
}

func (m *MockRepository) Get(key string) (string, bool) {
	val, exists := m.data[key]
	return val, exists
}

func setupRouter(handler *Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", handler.CreateShortURLHandler)
	r.Get("/{shortURL}", handler.GetOriginalURLHandler)
	return r
}

func TestCreateShortURLHandler(t *testing.T) {
	// Создаём мок-хранилище и сервис
	mockRepo := NewMockRepository()
	service := services.NewShortenerService(mockRepo)
	handler := NewHandler(service)
	router := setupRouter(handler)

	// Создаём HTTP-запрос
	body := strings.NewReader("https://practicum.yandex.ru/")
	req := httptest.NewRequest(http.MethodPost, "/", body)
	w := httptest.NewRecorder()

	// Выполняем запрос
	router.ServeHTTP(w, req)

	// Проверяем результат
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	responseBody := w.Body.String()
	if !strings.Contains(responseBody, "http://localhost:8080/") {
		t.Errorf("expected response to contain base URL, got %q", responseBody)
	}
}

func TestGetOriginalURLHandler(t *testing.T) {
	// Создаём мок-хранилище и сервис
	mockRepo := NewMockRepository()
	service := services.NewShortenerService(mockRepo)
	handler := NewHandler(service)
	router := setupRouter(handler)

	// Добавляем значение в хранилище
	shortURL := service.GenerateShortURL("https://practicum.yandex.ru/")

	// Случай, когда ключ существует
	req := httptest.NewRequest(http.MethodGet, "/"+shortURL, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
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
	router.ServeHTTP(w, req)

	resp = w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
