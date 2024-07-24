package database

import (
	"time"

	"github.com/GosMachine/ProductService/internal/models"
)

type Coupon interface {
	GetCoupon(code string) (models.Coupon, error)
	UpdateCoupon(coupon models.Coupon) error
	CreateCoupon(code string, value float64, valueType models.ValueTypeEnum, maxUsageCount, usageCount int, expiresAt *time.Time) error
	DeleteCoupon(code string) error
}

func (d *Database) GetCoupon(code string) (models.Coupon, error) {
	var coupon models.Coupon
	if err := d.db.Where("code = ?", code).First(&coupon).Error; err != nil {
		return models.Coupon{}, err
	}
	return coupon, nil
}

func (d *Database) UpdateCoupon(coupon models.Coupon) error {
	err := d.db.Save(&coupon).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateCoupon(code string, value float64, valueType models.ValueTypeEnum, maxUsageCount, usageCount int, expiresAt *time.Time) error {
	coupon := models.Coupon{
		Code:          code,
		Value:         value,
		ValueType:     valueType,
		MaxUsageCount: maxUsageCount,
		UsageCount:    usageCount,
		ExpiresAt:     expiresAt,
	}
	err := d.db.Create(&coupon).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteCoupon(code string) error {
	err := d.db.Where("code = ?", code).Delete(&models.Coupon{}).Error
	if err != nil {
		return err
	}
	return nil
}
