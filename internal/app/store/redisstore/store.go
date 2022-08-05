package redisstore

import (
	"github.com/go-redis/redis"
	"rest-redis-app/internal/app/store"
)

type Store struct {
	db              *redis.Client
	redisRepository *RedisRepository
}

func New(db *redis.Client) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Repository() store.Repository {
	if s.redisRepository != nil {
		return s.redisRepository
	}

	s.redisRepository = &RedisRepository{
		store: s,
	}

	return s.redisRepository
}

func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

//func (s *Store) Dec(key string) (interface{}, interface{}) {
//
//	oldValue, err := s.db.Get(key).Int()
//
//	if err != nil {
//		return "", err
//	}
//
//	_, err = s.db.Set(key, oldValue-1, 0).Result()
//
//	if err != nil {
//		return "", err
//	}
//
//	return s.db.Get(key).Val(), nil
//}
