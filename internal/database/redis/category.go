package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GosMachine/ProductService/internal/database/postgres"
	"github.com/GosMachine/ProductService/internal/models"
)

func (r *Redis) GetCategory(slug string) (*models.Category, error) {
	var (
		err      error
		category *models.Category
	)
	categoryJson := r.Client.Get(context.Background(), "category:"+slug).Val()
	if categoryJson != "" {
		err = json.Unmarshal([]byte(categoryJson), category)
		if err == nil {
			return category, nil
		}
	}
	category, err = r.db.GetCategory(slug)
	if err != nil {
		return nil, err
	}
	go r.SetCategoryCache(slug, category)
	return category, nil
}

func (r *Redis) SetCategoryCache(slug string, category *models.Category) error {
	categoryJson, err := json.Marshal(category)
	if err != nil {
		return err
	}
	r.Client.Set(context.Background(), "category:"+slug, categoryJson, time.Hour*12)
	return nil
}

func (r *Redis) GetCategories() (*postgres.Categories, error) {
	var categories *postgres.Categories
	categoriesJson, err := r.Client.Get(context.Background(), "categories").Bytes()
	if len(categoriesJson) > 0 && err == nil {
		err = json.Unmarshal(categoriesJson, categories)
		if err == nil {
			return categories, nil
		}
	}
	categories, err = r.db.GetCategories()
	if err != nil {
		return nil, err
	}
	go r.SetCategoriesCache(categories)
	return categories, nil
}

func (r *Redis) SetCategoriesCache(categories *postgres.Categories) error {
	categoriesJson, err := json.Marshal(categories)
	if err != nil {
		return err
	}
	r.Client.Set(context.Background(), "categories", categoriesJson, time.Hour*12)
	return nil
}
