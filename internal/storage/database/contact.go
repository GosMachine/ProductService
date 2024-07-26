package database

import "github.com/GosMachine/ProductService/internal/models"

type Contact interface {
	GetTicket(id int64) (models.Contact, error)
	CreateTicket(name, email, message, ip string) error
}

func (d *database) GetTicket(id int64) (models.Contact, error) {
	var contact models.Contact
	if err := d.db.Where("id = ?", id).First(&contact).Error; err != nil {
		return models.Contact{}, err
	}
	return contact, nil
}

func (d *database) CreateTicket(name, email, message, ip string) error {
	contact := models.Contact{Name: name, Email: email, Message: message, IP: ip}
	if err := d.db.Create(&contact).Error; err != nil {
		return err
	}
	return nil
}
