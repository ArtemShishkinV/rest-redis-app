package redisstore

import (
	"rest-redis-app/internal/app/store"
	"strconv"
	"testing"
)

var (
	keySuccessTest   = "test"
	valueSuccessTest = "10"
	valueErrorTest   = store.ErrRecordNotFound
)

var testStore *Store

func TestRedisRepository_FindValueSuccess(t *testing.T) {
	expected, err := strconv.Atoi(valueSuccessTest)
	if err != nil {
		t.Fatal("Expected value must be int!")
	}

	actual, err := testStore.Repository().FindIntValue(keySuccessTest)

	if err != nil {
		t.Fatal("Find value drop error...")
	}

	if actual != expected {
		t.Fatal("Actual value not equal expected!")
	}
}

func TestRedisRepository_FindValueError(t *testing.T) {
	expected := valueErrorTest

	_, err := testStore.Repository().FindIntValue("FAKE")
	if err != expected {
		t.Fatal("Test must be return error!")
	}
}

func TestRedisRepository_UpdateValueSuccess(t *testing.T) {
	expected := 11

	if err := testStore.Repository().UpdateValue(keySuccessTest, expected); err != nil {
		t.Fatal("Change value not success...")
	}

	actual, err := testStore.Repository().FindIntValue(keySuccessTest)

	if err != nil {
		t.Fatal("Find value drop error...")
	}

	if expected != actual {
		t.Fatal("Actual value not equal expected!")
	}
}
