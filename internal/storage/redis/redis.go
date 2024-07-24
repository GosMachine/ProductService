package redis

import (
	"github.com/GosMachine/ProductService/internal/storage/database"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	client *redis.Client
	db     *database.Database
	log    *zap.Logger
}

func New(db *database.Database, log *zap.Logger) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &Redis{client: client, db: db, log: log}
}
