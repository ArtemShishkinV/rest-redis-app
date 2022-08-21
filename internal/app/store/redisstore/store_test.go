package redisstore

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	s, _ := miniredis.Run()
	_ = s.Set(keySuccessTest, valueSuccessTest)

	if testStore = configureStore(s.Addr()); testStore == nil {
		os.Exit(1)
	}

	code := m.Run()

	defer testStore.Close()

	os.Exit(code)
}

func TestStore_Close(t *testing.T) {
	if err := testStore.Close(); err != nil {
		t.Fatal(err)
	}
	if err := testStore.Close(); err == nil {
		t.Fatal(err)
	}
}

func configureStore(databaseURL string) *Store {
	db := redis.NewClient(&redis.Options{
		Addr: databaseURL,
	})

	if err := db.Ping().Err(); err != nil {
		return nil
	}

	return New(db)
}
