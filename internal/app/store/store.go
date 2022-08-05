package store

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Store struct {
	config *Config
	db     *redis.Client
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db := redis.NewClient(&redis.Options{
		Addr:     s.config.Host + s.config.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := db.Ping().Result()

	if err != nil {
		return err
	}

	fmt.Println(db.Get("key").String())

	s.db = db

	return nil
}

func (s *Store) Inc(key string) (string, error) {

	oldValue, err := s.db.Get(key).Int()

	if err != nil {
		return "", err
	}

	_, err = s.db.Set(key, oldValue+1, 0).Result()

	if err != nil {
		return "", err
	}

	return s.db.Get(key).Val(), nil
}

func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Dec(key string) (interface{}, interface{}) {

	oldValue, err := s.db.Get(key).Int()

	if err != nil {
		return "", err
	}

	_, err = s.db.Set(key, oldValue-1, 0).Result()

	if err != nil {
		return "", err
	}

	return s.db.Get(key).Val(), nil
}
