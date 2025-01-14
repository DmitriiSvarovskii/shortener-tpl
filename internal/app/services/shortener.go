package services

import (
	"errors"
	"math/rand"

	"github.com/DmitriiSvarovskii/shortener-tpl.git/internal/app/storage"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const lenURL = 8

// ErrKeyNotFound возвращается, если ключ отсутствует в хранилище.
var ErrKeyNotFound = errors.New("key not found")

type ShortenerService struct {
	repo storage.Repository
}

func NewShortenerService(repo storage.Repository) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) GenerateShortURL(value string) string {
	key := randStr(lenURL)
	s.repo.Save(key, value)
	return key
}

func (s *ShortenerService) GetOriginalURL(key string) (string, error) {
	val, exists := s.repo.Get(key)
	if !exists {
		return "", ErrKeyNotFound
	}
	return val, nil
}

func randStr(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(buf)
}
