package services

import (
	"testing"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/storage"
)

func TestGenerateShortURL(t *testing.T) {
	// Создаём моковое хранилище
	mockRepo := storage.NewMemoryRepository()
	service := NewShortenerService(mockRepo)

	originalURL := "https://practicum.yandex.ru/"
	shortURL := service.GenerateShortURL(originalURL)

	// Проверяем, что короткий URL не пустой
	if shortURL == "" {
		t.Errorf("expected non-empty short URL, got empty")
	}

	// Проверяем, что оригинальный URL сохраняется в хранилище
	savedURL, exists := mockRepo.Get(shortURL)
	if !exists {
		t.Errorf("expected URL to be saved in repository, but it was not")
	}
	if savedURL != originalURL {
		t.Errorf("expected saved URL %q, got %q", originalURL, savedURL)
	}
}

func TestGetOriginalURL(t *testing.T) {
	// Создаём моковое хранилище
	mockRepo := storage.NewMemoryRepository()
	service := NewShortenerService(mockRepo)

	originalURL := "https://practicum.yandex.ru/"
	shortURL := "short123"

	// Сохраняем URL вручную
	mockRepo.Save(shortURL, originalURL)

	// Тестируем успешное получение
	returnedURL, err := service.GetOriginalURL(shortURL)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if returnedURL != originalURL {
		t.Errorf("expected URL %q, got %q", originalURL, returnedURL)
	}

	// Тестируем случай, когда ключ не найден
	_, err = service.GetOriginalURL("nonexistentKey")
	if err != ErrKeyNotFound {
		t.Errorf("expected error %v, got %v", ErrKeyNotFound, err)
	}
}