package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

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
	ID            int `gorm:"primary_key"`
	Name          string
	Slug          string `gorm:"unique"`
	Description   string
	Price         float64
	Stock         int64
	CategoryID    int
	NumberOfSales int
	Category      Category
	Fields        InputFields
	ImageURL      string
}

type InputField struct {
	Label string
	Type  string
}

type InputFields []InputField

//todo product review

func (sla *InputFields) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &sla)
}

func (sla InputFields) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}
