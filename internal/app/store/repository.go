package store

type Repository interface {
	FindIntValue(key string) (int, error)
	UpdateValue(key string, val int) error
}
