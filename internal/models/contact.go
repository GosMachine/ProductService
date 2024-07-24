package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID      int `gorm:"primary_key"`
	Name    string
	Email   string
	Message string
	IP      string
}
