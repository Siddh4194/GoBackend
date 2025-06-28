package models

import "time"

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"not null"`
	TotalAmount float64
	Status      string `gorm:"default:PENDING"` // or ENUM type in PostgreSQL
	CreatedAt   time.Time

	User       User
	OrderItems []OrderItem
}

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	ProductID  uint    `gorm:"not null"`
	QuantityKG float64 `gorm:"not null"`
	PricePerKG float64 `gorm:"not null"`

	Order   Order
	Product Product
}
