package storage

type MemoryRepository struct {
	data map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]string),
	}
}

func (m *MemoryRepository) Save(key, value string) {
	m.data[key] = value
}

func (m *MemoryRepository) Get(key string) (string, bool) {
	val, ok := m.data[key]
	return val, ok
}
