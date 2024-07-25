package database

import (
	"github.com/GosMachine/ProductService/internal/models"
)

type Product interface {
	CreateCategory(name, description string) (models.Category, error)
	GetCategory(slug string) (models.Category, error)
	GetCategories() ([]Category, error)
}

type Category struct {
	Name string
	Slug string
}

func (d *Database) CreateCategory(name, slug, description string) (models.Category, error) {
	category := models.Category{Name: name, Slug: slug, Description: description}

	if err := d.db.Create(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (d *Database) GetCategory(slug string) (models.Category, error) {
	var category models.Category
	if err := d.db.Preload("Products").Where("slug = ?", slug).First(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (d *Database) GetCategories() ([]Category, error) {
	var (
		categoriesFromDb []*models.Category
		categories       []Category
	)
	if err := d.db.Select("name", "slug").Find(&categoriesFromDb).Error; err != nil {
		return []Category{}, err
	}
	for _, category := range categoriesFromDb {
		categories = append(categories, Category{Name: category.Name, Slug: category.Slug})
	}
	return categories, nil
}
