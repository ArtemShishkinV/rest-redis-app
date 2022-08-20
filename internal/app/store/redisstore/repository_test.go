package redisstore

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"strconv"
	"testing"
)

var (
	keyTest   = "test"
	valueTest = "10"
)

func setupStore() *Store {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	if err = s.Set(keyTest, valueTest); err != nil {
		panic(err)
	}

	db := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	return New(db)
}

func TestRedisRepository_FindValueSuccess(t *testing.T) {
	store := setupStore()

	defer store.Close()

	expected, _ := strconv.Atoi(valueTest)
	actual, _ := store.Repository().FindValue(keyTest)

	if actual != expected {
		t.Fatal("Actual value not equal expected!")
	}
}
