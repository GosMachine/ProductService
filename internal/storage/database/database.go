package database

import (
	"fmt"
	"os"

	"github.com/GosMachine/ProductService/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

type Database interface {
	Product
	Contact
	Coupon
}

func New() (Database, error) {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models.Category{}, models.Coupon{}, models.Product{}, models.Contact{})
	if err != nil {
		return nil, err
	}
	return &database{db: db}, nil
}
