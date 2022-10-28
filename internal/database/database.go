package database

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	client *redis.Client
}

func New(address, password string) (*Database, error) {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
	}, nil
}

func (db *Database) Close() error {
	return db.client.Close()
}

func (db *Database) marshal(v interface{}) (string, error) {
	m, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(m), nil
}

func (db *Database) unmarshal(m string, v interface{}) error {
	return json.Unmarshal([]byte(m), v)
}
