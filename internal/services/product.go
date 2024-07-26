package services

import (
	"errors"
	"fmt"

	"github.com/GosMachine/ProductService/internal/models"
	"github.com/GosMachine/ProductService/internal/storage/database"
	"github.com/GosMachine/ProductService/internal/storage/redis"
	"go.uber.org/zap"
)

var ErrCategoryNotFound = errors.New("category not found")

type Product struct {
	log   *zap.Logger
	db    *database.Database
	redis redis.Service
}

func New(log *zap.Logger, db *database.Database, redis redis.Service) *Product {
	return &Product{
		log:   log,
		db:    db,
		redis: redis,
	}
}

func (p *Product) GetCategory(slug string) (*models.Category, error) {
	log := p.log.With(
		zap.String("slug", slug),
	)
	log.Info("getting category")
	category, err := p.redis.GetCategory(slug)
	if err != nil {
		log.Error("failed to get category", zap.Error(err))
		return nil, ErrCategoryNotFound
	}
	log.Info("category successfully taken")
	return category, nil
}

func (p *Product) GetCategories() ([]database.Category, error) {
	p.log.Info("getting categories")
	categories, err := p.redis.GetCategories()
	if err != nil {
		p.log.Error("failed to get categories", zap.Error(err))
		return nil, ErrCategoryNotFound
	}
	p.log.Info("categories successfully taken")
	return categories, nil
}

func (p *Product) CreateTicket(name, email, message, ip string) error {
	p.log.Info("creating ticket", zap.String("email", email))
	err := p.db.CreateTicket(name, email, message, ip)
	if err != nil {
		p.log.Error("err create ticket", zap.Error(err))
		return fmt.Errorf("err create ticket")
	}
	return nil
}

func (p *Product) GetProduct(categorySlug, productSlug string) (product *models.Product, err error) {
	p.log.Info("getting product", zap.String("category", categorySlug), zap.String("product", productSlug))
	category, err := p.redis.GetCategory(categorySlug)
	if err != nil {
		p.log.Error("failed to get category", zap.String("category", categorySlug), zap.Error(err))
		return nil, ErrCategoryNotFound
	}
	for _, product := range category.Products {
		if product.Slug == productSlug {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}
