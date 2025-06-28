package migrations

import (
	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Ad{},
	)
}
