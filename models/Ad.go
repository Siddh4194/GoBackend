package models

import "time"

type Ad struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255;not null"`
	ImageURL  string
	Link      string
	Active    bool `gorm:"default:true"`
	CreatedAt time.Time
}
