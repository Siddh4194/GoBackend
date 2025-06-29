package models

import "time"

type CartItem struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index;not null"`
	ProductID  uint      `gorm:"index;not null"`
	QuantityKG float64   `gorm:"not null"`
	AddedAt    time.Time `gorm:"autoCreateTime"`

	User    User
	Product Product

	// Enforce unique(user_id, product_id)
}

func (CartItem) TableName() string {
	return "cart_items"
}
