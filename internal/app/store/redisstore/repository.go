package redisstore

type RedisRepository struct {
	store *Store
}

func (r *RedisRepository) IncrementKeyByValue(key string, val int) (int, error) {

	oldValue, err := r.store.db.Get(key).Int()

	if err != nil {
		return oldValue, err
	}

	newValue := oldValue + val
	_, err = r.store.db.Set(key, newValue, 0).Result()

	if err != nil {
		return oldValue, err
	}

	return newValue, nil
}
