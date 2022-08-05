package apiserver

import (
	"github.com/go-redis/redis"
	"net/http"
	"rest-redis-app/internal/app/store/redisstore"
)

func Start(config *Config) error {
	db, err := newDB(config.Store.Host, config.Store.Port)
	if err != nil {
		return err
	}

	defer db.Close()

	store := redisstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(host string, port string) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     host + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := db.Ping().Result()

	if err != nil {
		return nil, err
	}

	return db, nil
}
