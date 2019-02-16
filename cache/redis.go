package cache

import (
	"github.com/go-redis/redis"
	"github.com/ttacon/chalk"
	"log"
	"time"
)

// RedisTore is a struct that holds redis client for easy querying later. This concrete type will implement the Cache
// interface
type RedisStore struct {
	client *redis.Client
}

// Get accepts a key and returns the value from the redis data store as byte slice.
// It also returns an error
func (m *RedisStore) Get(key string) ([]byte, error) {
	val, err := m.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// Set accepts key, value (content) and duration for the data. Then saves it to the redis store
func (m *RedisStore) Set(key string, content []byte, duration string) {
	d, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = m.client.Set(key, content, d).Result()
}

// NewMemoryStorage is the factory pattern function that accepts address, password and db as argument and
// returns a new RedisStore object as Cache interface. It also returns an error
func NewMemoryStorage(addr, password string, db int) (Cache, error) {
	log.Println(chalk.Green, "redis: connecting....")
	client := redis.NewClient(&redis.Options{
		Addr:addr,
		Password:password,
		DB: db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	log.Println(chalk.Green, "redis: connected")
	return &RedisStore{client: client}, nil
}


