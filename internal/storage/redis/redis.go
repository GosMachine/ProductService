package redis

import (
	"os"

	"github.com/GosMachine/ProductService/internal/models"
	"github.com/GosMachine/ProductService/internal/storage/database"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	client *redis.Client
	db     database.Database
	log    *zap.Logger
}

type Service interface {
	GetCategory(slug string) (category *models.Category, err error)
	GetCategories() (categories []database.Category, err error)
	SetCategoryCache(slug string, category *models.Category) error
	SetCategoriesCache(categories []database.Category) error
}

func New(db database.Database, log *zap.Logger) Service {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	return &Redis{client: client, db: db, log: log}
}
