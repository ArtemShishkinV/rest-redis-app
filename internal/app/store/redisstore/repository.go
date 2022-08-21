package redisstore

import "rest-redis-app/internal/app/store"

type RedisRepository struct {
	store *Store
}

func (r *RedisRepository) FindIntValue(key string) (int, error) {
	value, err := r.store.db.Get(key).Int()

	if err != nil {
		return 0, store.ErrRecordNotFound
	}

	return value, nil
}

func (r *RedisRepository) UpdateValue(key string, value int) error {
	_, err := r.store.db.Set(key, value, 0).Result()

	return err
}
