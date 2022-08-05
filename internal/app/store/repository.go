package store

type Repository interface {
	IncrementKeyByValue(key string, val int) (int, error)
}
