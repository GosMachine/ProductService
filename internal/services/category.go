package services

import (
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
