package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Name        string
	Slug        string `gorm:"unique"`
	Description string
	Products    []Product
}

type Product struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	Name        string
	Slug        string `gorm:"unique"`
	Description string
	Price       float64
	Stock       int64
	CategoryID  int
	Category    Category
	Fields      InputFields
	ImageURL    string
}

type InputFields struct {
	Label string
	Type  string
}

//todo product review
