package services

import (
	"errors"

	"github.com/GosMachine/ProductService/internal/database/postgres"
	"github.com/GosMachine/ProductService/internal/database/redis"
	"github.com/GosMachine/ProductService/internal/models"
	"go.uber.org/zap"
)

var ErrCategoryNotFound = errors.New("category not found")

type Product struct {
	log   *zap.Logger
	db    Database
	redis Redis
}

type Database interface {
}

type Redis interface {
	GetCategory(name string) (category *models.Category, err error)
}

func New(log *zap.Logger, db *postgres.Storage, redis *redis.Redis) *Product {
	return &Product{
		log:   log,
		db:    db,
		redis: redis,
	}
}
