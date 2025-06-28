package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Phone     string    `gorm:"size:15;unique;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	Password  string    `gorm:"not null"`
	Address   string
	CreatedAt time.Time

	CartItems []CartItem
	Orders    []Order
}

