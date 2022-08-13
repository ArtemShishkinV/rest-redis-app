package store

type Repository interface {
	FindValue(key string) (int, error)
	UpdateValue(key string, val int) error
}
