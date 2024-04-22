package postgres

import (
	"github.com/GosMachine/ProductService/internal/models"
)

type Categories struct {
	Names []string
	Slugs []string
}

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

func (s *Storage) GetCategories() (*Categories, error) {
	var (
		categoriesFromDb []*models.Category
		categories       Categories
	)
	if err := s.db.Select("name", "slug").Find(&categoriesFromDb).Error; err != nil {
		return nil, err
	}
	for _, category := range categoriesFromDb {
		categories.Names = append(categories.Names, category.Name)
		categories.Slugs = append(categories.Slugs, category.Slug)
	}
	return &categories, nil
}
