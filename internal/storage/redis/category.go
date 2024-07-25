package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GosMachine/ProductService/internal/models"
	"github.com/GosMachine/ProductService/internal/storage/database"
	"go.uber.org/zap"
)

func (r *Redis) GetCategory(slug string) (*models.Category, error) {
	var (
		err      error
		category models.Category
	)
	categoryJson := r.client.Get(context.Background(), "category:"+slug).Val()
	if categoryJson != "" {
		err := json.Unmarshal([]byte(categoryJson), &category)
		if err == nil {
			return &category, nil
		}
	}
	category, err = r.db.GetCategory(slug)
	if err != nil {
		return nil, err
	}
	go func() {
		err := r.SetCategoryCache(slug, &category)
		if err != nil {
			r.log.Error("err set category cache", zap.Error(err))
		}
	}()
	return &category, nil
}

func (r *Redis) SetCategoryCache(slug string, category *models.Category) error {
	categoryJson, err := json.Marshal(category)
	if err != nil {
		return err
	}
	r.client.Set(context.Background(), "category:"+slug, categoryJson, time.Hour*12)
	return nil
}

func (r *Redis) GetCategories() ([]database.Category, error) {
	var categories []database.Category
	categoriesJson, err := r.client.Get(context.Background(), "categories").Bytes()
	if len(categoriesJson) > 0 && err == nil {
		err = json.Unmarshal(categoriesJson, &categories)
		if err == nil {
			return categories, nil
		}
	}
	categories, err = r.db.Product.GetCategories()
	if err != nil {
		return nil, err
	}
	go func() {
		err := r.SetCategoriesCache(categories)
		if err != nil {
			r.log.Error("err set categories cache", zap.Error(err))
		}
	}()
	return categories, nil
}

func (r *Redis) SetCategoriesCache(categories []database.Category) error {
	categoriesJson, err := json.Marshal(categories)
	if err != nil {
		return err
	}
	r.client.Set(context.Background(), "categories", categoriesJson, time.Hour*12)
	return nil
}
