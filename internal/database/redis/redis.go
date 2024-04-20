package redis

import (
	"github.com/GosMachine/ProductService/internal/database/postgres"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
	db     *postgres.Storage
}

func New(db *postgres.Storage) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &Redis{Client: client, db: db}
}
