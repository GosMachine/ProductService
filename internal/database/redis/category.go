package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GosMachine/ProductService/internal/models"
)

func (r *Redis) GetCategory(name string) (*models.Category, error) {
	var (
		err      error
		category *models.Category
	)
	categoryJson := r.Client.Get(context.Background(), "category:"+name).Val()
	if categoryJson != "" {
		err = json.Unmarshal([]byte(categoryJson), category)
		if err == nil {
			return category, nil
		}
	}
	category, err = r.db.GetCategory(name)
	if err != nil {
		return nil, err
	}
	go r.SetCategoryCache(name, category)
	return category, nil
}

func (r *Redis) SetCategoryCache(name string, category *models.Category) error {
	categoryJson, err := json.Marshal(category)
	if err != nil {
		return err
	}
	r.Client.Set(context.Background(), "category:"+name, categoryJson, time.Hour*12)
	return nil
}

func (r *Redis) GetCategoryNames() ([]string, error) {
	categories, err := r.Client.LRange(context.Background(), "categoryNames", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	if len(categories) >= 1 {
		return categories, nil
	}
	categories, err = r.db.GetCategoryNames()
	if err != nil {
		return nil, err
	}
	go r.SetCategoryNamesCache(categories)
	return categories, nil
}

func (r *Redis) SetCategoryNamesCache(categories []string) error {
	for _, v := range categories {
		err := r.Client.RPush(context.Background(), "categoryNames", v).Err()
		if err != nil {
			return err
		}
	}
	err := r.Client.Expire(context.Background(), "categoryNames", time.Hour*12).Err()
	if err != nil {
		return err
	}
	return nil
}
