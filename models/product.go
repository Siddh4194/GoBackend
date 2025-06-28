package models

import "time"

type Product struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"size:100;not null"`
	Description     string
	PricePerKG      float64 `gorm:"not null"`
	ImageURL        string
	StockQuantityKG float64
	CategoryID      *uint
	Category        Category `gorm:"foreignKey:CategoryID"`
	CreatedAt       time.Time

	CartItems  []CartItem
	OrderItems []OrderItem
}
