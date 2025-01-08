package storage

import (
	"testing"
)

func TestMemoryRepository_SaveAndGet(t *testing.T) {
	repo := NewMemoryRepository()

	// Тестируем сохранение ключа и значения
	key := "short123"
	value := "https://practicum.yandex.ru/"
	repo.Save(key, value)

	// Проверяем, что значение можно получить
	returnedValue, exists := repo.Get(key)
	if !exists {
		t.Errorf("expected key %q to exist, but it does not", key)
	}
	if returnedValue != value {
		t.Errorf("expected value %q, got %q", value, returnedValue)
	}

	// Тестируем случай, когда ключ не существует
	_, exists = repo.Get("nonexistentKey")
	if exists {
		t.Errorf("expected key %q to not exist, but it does", "nonexistentKey")
	}
}
