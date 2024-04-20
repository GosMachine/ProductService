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
	r.SetCategoryCache(name, category)
	return category, nil
}

func (r *Redis) SetCategoryCache(name string, category *models.Category) error {
	categoryJson, err := json.Marshal(category)
	if err != nil {
		return err
	}
	r.Client.Set(context.Background(), "category:"+name, categoryJson, time.Hour*24)
	return nil
}
