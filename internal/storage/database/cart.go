package database

import (
	"github.com/GosMachine/ProductService/internal/models"
)

type Cart interface {
	DeleteItem(id string) error
}

func (d *Database) DeleteItem(id string) error {
	err := d.db.Where("id = ?", id).Delete(&models.CartItem{}).Error
	if err != nil {
		return err
	}
	return nil
}
