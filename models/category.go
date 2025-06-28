package models

import "time"

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;unique;not null"`
	Description string
	CreatedAt   time.Time

	Products []Product
}
