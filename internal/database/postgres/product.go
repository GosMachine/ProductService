package postgres

import (
	"github.com/GosMachine/ProductService/internal/models"
)

func (s *Storage) CreateCategory(name, description string) (models.Category, error) {
	category := models.Category{Name: name, Description: description, Products: []models.Product{}}

	if err := s.db.Create(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (s *Storage) GetCategory(name string) (*models.Category, error) {
	var category models.Category
	if err := s.db.Preload("Products").Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
