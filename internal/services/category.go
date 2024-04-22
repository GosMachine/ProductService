package services

import (
	"github.com/GosMachine/ProductService/internal/database/postgres"
	"github.com/GosMachine/ProductService/internal/models"
	"go.uber.org/zap"
)

func (p *Product) GetCategory(name string) (*models.Category, error) {
	log := p.log.With(
		zap.String("name", name),
	)
	log.Info("getting category")
	category, err := p.redis.GetCategory(name)
	if err != nil {
		log.Error("failed to get category", zap.Error(err))
		return nil, ErrCategoryNotFound
	}
	return category, nil
}

func (p *Product) GetCategories() (*postgres.Categories, error) {
	p.log.Info("getting categories")
	categories, err := p.redis.GetCategories()
	if err != nil {
		p.log.Error("failed to get categories", zap.Error(err))
		return nil, ErrCategoryNotFound
	}
	return categories, nil
}
