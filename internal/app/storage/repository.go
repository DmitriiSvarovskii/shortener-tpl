package storage

type Repository interface {
	Save(key, value string)
	Get(key string) (string, bool)
}
